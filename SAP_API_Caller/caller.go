package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-project-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetProject(projectInternalID, wBSElementInternalID, plant string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Project":
			func() {
				c.Project(projectInternalID, plant)
				wg.Done()
			}()

		case "WBSElement":
			func() {
				c.WBSElement(wBSElementInternalID, plant)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Project(projectInternalID, plant string) {
	projectData, err := c.callProjectSrvAPIRequirementProject("Project", projectInternalID, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(projectData)

	wBSElementData, err := c.callToWBSElement(projectData[0].ToWBSElement)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(wBSElementData)
}

func (c *SAPAPICaller) callProjectSrvAPIRequirementProject(api, projectInternalID, plant string) ([]sap_api_output_formatter.Project, error) {
	url := strings.Join([]string{c.baseURL, "API_PROJECT_SRV", api}, "/")
	param := c.getQueryWithProject(map[string]string{}, projectInternalID, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToProject(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToWBSElement(url string) ([]sap_api_output_formatter.ToWBSElement, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToWBSElement(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) WBSElement(wBSElementInternalID, plant string) {
	data, err := c.callProjectSrvAPIRequirementWBSElement("WBSElement", wBSElementInternalID, plant)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callProjectSrvAPIRequirementWBSElement(api, wBSElementInternalID, plant string) ([]sap_api_output_formatter.WBSElement, error) {
	url := strings.Join([]string{c.baseURL, "API_PROJECT_SRV", api}, "/")

	param := c.getQueryWithWBSElement(map[string]string{}, wBSElementInternalID, plant)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToWBSElement(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithProject(params map[string]string, projectInternalID, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("ProjectInternalID eq '%s' and Plant eq '%s'", projectInternalID, plant)
	return params
}

func (c *SAPAPICaller) getQueryWithWBSElement(params map[string]string, wBSElementInternalID, plant string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("WBSElementInternalID eq '%s' and Plant eq '%s'", wBSElementInternalID, plant)
	return params
}
