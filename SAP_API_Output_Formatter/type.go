package sap_api_output_formatter

type Project struct {
	ProjectInternalID             string `json:"ProjectInternalID"`
	ProjectExternalID             string `json:"ProjectExternalID"`
	ProjectDescription            string `json:"ProjectDescription"`
	ProjectLangBsdDescription     string `json:"ProjectLangBsdDescription"`
	ProjectProfileCode            string `json:"ProjectProfileCode"`
	CompanyCode                   string `json:"CompanyCode"`
	ControllingArea               string `json:"ControllingArea"`
	FunctionalArea                string `json:"FunctionalArea"`
	ProfitCenter                  string `json:"ProfitCenter"`
	PlannedStartDate              string `json:"PlannedStartDate"`
	PlannedEndDate                string `json:"PlannedEndDate"`
	WorkCenterLocation            string `json:"WorkCenterLocation"`
	ResponsiblePerson             string `json:"ResponsiblePerson"`
	ResponsiblePersonName         string `json:"ResponsiblePersonName"`
	ApplicantCode                 string `json:"ApplicantCode"`
	ApplicantName                 string `json:"ApplicantName"`
	CreationDate                  string `json:"CreationDate"`
	LastChangeDate                string `json:"LastChangeDate"`
	BasicDatesLastScheduledDate   string `json:"BasicDatesLastScheduledDate"`
	FcstdDatesLastScheduledDate   string `json:"FcstdDatesLastScheduledDate"`
	FactoryCalendar               string `json:"FactoryCalendar"`
	SchedulingDurationUnit        string `json:"SchedulingDurationUnit"`
	ForecastedStartDate           string `json:"ForecastedStartDate"`
	ForecastedEndDate             string `json:"ForecastedEndDate"`
	BusinessArea                  string `json:"BusinessArea"`
	Plant                         string `json:"Plant"`
	Currency                      string `json:"Currency"`
	BudgetProfile                 string `json:"BudgetProfile"`
	PlanningProfile               string `json:"PlanningProfile"`
	InvestmentProfile             string `json:"InvestmentProfile"`
	ProjInterestCalcProfile       string `json:"ProjInterestCalcProfile"`
	ResultAnalysisInternalID      string `json:"ResultAnalysisInternalID"`
	NetworkProfile                string `json:"NetworkProfile"`
	WBSSchedulingProfile          string `json:"WBSSchedulingProfile"`
	PlanningMethForProjBasicDate  string `json:"PlanningMethForProjBasicDate"`
	PlanningMethForProjFcstdDate  string `json:"PlanningMethForProjFcstdDate"`
	NetworkAssignmentType         string `json:"NetworkAssignmentType"`
	WBSIsStatisticalWBSElement    bool   `json:"WBSIsStatisticalWBSElement"`
	WBSIsMarkedForIntegratedPlng  bool   `json:"WBSIsMarkedForIntegratedPlng"`
	ProjectHasOwnStock            bool   `json:"ProjectHasOwnStock"`
	InventorySpecialStockValnType string `json:"InventorySpecialStockValnType"`
	WBSIsMarkedForAutomReqmtGrpg  bool   `json:"WBSIsMarkedForAutomReqmtGrpg"`
	SalesOrganization             string `json:"SalesOrganization"`
	DistributionChannel           string `json:"DistributionChannel"`
	Language                      string `json:"Language"`
	WBSElementExternalID          string `json:"WBSElementExternalID"`
	Division                      string `json:"Division"`
	JointVenture                  string `json:"JointVenture"`
	JointVentureEquityType        string `json:"JointVentureEquityType"`
	JointVentureObjectType        string `json:"JointVentureObjectType"`
	StatusProfile                 string `json:"StatusProfile"`
	WBSStatusProfile              string `json:"WBSStatusProfile"`
	SimulationProfile             string `json:"SimulationProfile"`
	SchedulingScenario            string `json:"SchedulingScenario"`
	DistributionProfile           string `json:"DistributionProfile"`
	PartnerDeterminationProcedure string `json:"PartnerDeterminationProcedure"`
	FreeDefinedTableFieldSemantic string `json:"FreeDefinedTableFieldSemantic"`
	FreeDefinedAttribute01        string `json:"FreeDefinedAttribute01"`
	FreeDefinedAttribute02        string `json:"FreeDefinedAttribute02"`
	FreeDefinedAttribute03        string `json:"FreeDefinedAttribute03"`
	FreeDefinedAttribute04        string `json:"FreeDefinedAttribute04"`
	FreeDefinedQuantity1          string `json:"FreeDefinedQuantity1"`
	FreeDefinedQuantity1Unit      string `json:"FreeDefinedQuantity1Unit"`
	FreeDefinedQuantity2          string `json:"FreeDefinedQuantity2"`
	FreeDefinedQuantity2Unit      string `json:"FreeDefinedQuantity2Unit"`
	FreeDefinedAmount1            string `json:"FreeDefinedAmount1"`
	FreeDefinedAmount1Currency    string `json:"FreeDefinedAmount1Currency"`
	FreeDefinedAmount2            string `json:"FreeDefinedAmount2"`
	FreeDefinedAmount2Currency    string `json:"FreeDefinedAmount2Currency"`
	FreeDefinedDate1              string `json:"FreeDefinedDate1"`
	FreeDefinedDate2              string `json:"FreeDefinedDate2"`
	FreeDefinedIndicator1         bool   `json:"FreeDefinedIndicator1"`
	FreeDefinedIndicator2         bool   `json:"FreeDefinedIndicator2"`
	ToWBSElement                  string `json:"to_WBSElement"`
}

type ToWBSElement struct {
	WBSElementInternalID           string `json:"WBSElementInternalID"`
	WBSElementExternalID           string `json:"WBSElementExternalID"`
	WBSElementShortID              string `json:"WBSElementShortID"`
	WBSDescription                 string `json:"WBSDescription"`
	WBSElementLangBsdDescription   string `json:"WBSElementLangBsdDescription"`
	ResponsiblePerson              string `json:"ResponsiblePerson"`
	ResponsiblePersonName          string `json:"ResponsiblePersonName"`
	ApplicantCode                  string `json:"ApplicantCode"`
	ApplicantName                  string `json:"ApplicantName"`
	CompanyCode                    string `json:"CompanyCode"`
	BusinessArea                   string `json:"BusinessArea"`
	ControllingArea                string `json:"ControllingArea"`
	FunctionalArea                 string `json:"FunctionalArea"`
	ProfitCenter                   string `json:"ProfitCenter"`
	ResponsibleCostCenter          string `json:"ResponsibleCostCenter"`
	Plant                          string `json:"Plant"`
	FreeDefinedTableFieldSemantic  string `json:"FreeDefinedTableFieldSemantic"`
	FactoryCalendar                string `json:"FactoryCalendar"`
	PriorityCode                   string `json:"PriorityCode"`
	Currency                       string `json:"Currency"`
	CostingSheet                   string `json:"CostingSheet"`
	CostCenter                     string `json:"CostCenter"`
	RequestingCostCenter           string `json:"RequestingCostCenter"`
	ProjectInternalID              string `json:"ProjectInternalID"`
	WBSElementIsBillingElement     bool   `json:"WBSElementIsBillingElement"`
	InvestmentProfile              string `json:"InvestmentProfile"`
	WBSIsStatisticalWBSElement     bool   `json:"WBSIsStatisticalWBSElement"`
	WBSIsAccountAssignmentElement  bool   `json:"WBSIsAccountAssignmentElement"`
	ProjectType                    string `json:"ProjectType"`
	WBSElementIsPlanningElement    bool   `json:"WBSElementIsPlanningElement"`
	WorkCenterLocation             string `json:"WorkCenterLocation"`
	ResultAnalysisInternalID       string `json:"ResultAnalysisInternalID"`
	TaxJurisdiction                string `json:"TaxJurisdiction"`
	FunctionalLocation             string `json:"FunctionalLocation"`
	CreationDate                   string `json:"CreationDate"`
	LastChangeDate                 string `json:"LastChangeDate"`
	WBSIsMarkedForIntegratedPlng   bool   `json:"WBSIsMarkedForIntegratedPlng"`
	Equipment                      string `json:"Equipment"`
	ProjectObjectChangeNumber      string `json:"ProjectObjectChangeNumber"`
	WBSElementHierarchyLevel       int    `json:"WBSElementHierarchyLevel"`
	ReferenceElement               string `json:"ReferenceElement"`
	ProjInterestCalcProfile        string `json:"ProjInterestCalcProfile"`
	Language                       string `json:"Language"`
	IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
	BasicStartDate                 string `json:"BasicStartDate"`
	ForecastedStartDate            string `json:"ForecastedStartDate"`
	ActualStartDate                string `json:"ActualStartDate"`
	BasicEndDate                   string `json:"BasicEndDate"`
	ForecastedEndDate              string `json:"ForecastedEndDate"`
	ActualEndDate                  string `json:"ActualEndDate"`
	BasicDuration                  string `json:"BasicDuration"`
	BasicDurationUnit              string `json:"BasicDurationUnit"`
	ForecastedDuration             string `json:"ForecastedDuration"`
	ForecastedDurationUnit         string `json:"ForecastedDurationUnit"`
	ActualDuration                 string `json:"ActualDuration"`
	ActualDurationUnit             string `json:"ActualDurationUnit"`
	SchedldBasicEarliestStartDate  string `json:"SchedldBasicEarliestStartDate"`
	ScheduledBasicLatestEndDate    string `json:"ScheduledBasicLatestEndDate"`
	SchedldFcstdEarliestStartDate  string `json:"SchedldFcstdEarliestStartDate"`
	LatestSchedldFcstdEndDate      string `json:"LatestSchedldFcstdEndDate"`
	TentativeActualStartDate       string `json:"TentativeActualStartDate"`
	TentativeActualEndDate         string `json:"TentativeActualEndDate"`
	JointVenture                   string `json:"JointVenture"`
	JointVentureEquityType         string `json:"JointVentureEquityType"`
	JntVntrProjectType             string `json:"JntVntrProjectType"`
	SubProject                     string `json:"SubProject"`
	InvestmentReason               string `json:"InvestmentReason"`
	InvestmentScale                string `json:"InvestmentScale"`
	EnvironmentalInvestmentReason  string `json:"EnvironmentalInvestmentReason"`
	RequestingCompanyCode          string `json:"RequestingCompanyCode"`
	NetworkAssignmentType          string `json:"NetworkAssignmentType"`
	CostObject                     string `json:"CostObject"`
	WBSElementParentInternalID     string `json:"WBSElementParentInternalID"`
	WBSElementChildInternalID      string `json:"WBSElementChildInternalID"`
	LeftSiblingWBSElmntInternalID  string `json:"LeftSiblingWBSElmntInternalID"`
	RightSiblingWBSElmntInternalID string `json:"RightSiblingWBSElmntInternalID"`
	FreeDefinedAttribute01         string `json:"FreeDefinedAttribute01"`
	FreeDefinedAttribute02         string `json:"FreeDefinedAttribute02"`
	FreeDefinedAttribute03         string `json:"FreeDefinedAttribute03"`
	FreeDefinedAttribute04         string `json:"FreeDefinedAttribute04"`
	FreeDefinedQuantity1           string `json:"FreeDefinedQuantity1"`
	FreeDefinedQuantity1Unit       string `json:"FreeDefinedQuantity1Unit"`
	FreeDefinedQuantity2           string `json:"FreeDefinedQuantity2"`
	FreeDefinedQuantity2Unit       string `json:"FreeDefinedQuantity2Unit"`
	FreeDefinedAmount1             string `json:"FreeDefinedAmount1"`
	FreeDefinedAmount1Currency     string `json:"FreeDefinedAmount1Currency"`
	FreeDefinedAmount2             string `json:"FreeDefinedAmount2"`
	FreeDefinedAmount2Currency     string `json:"FreeDefinedAmount2Currency"`
	FreeDefinedDate1               string `json:"FreeDefinedDate1"`
	FreeDefinedDate2               string `json:"FreeDefinedDate2"`
	FreeDefinedIndicator1          bool   `json:"FreeDefinedIndicator1"`
	FreeDefinedIndicator2          bool   `json:"FreeDefinedIndicator2"`
}

type WBSElement struct {
	WBSElementInternalID           string `json:"WBSElementInternalID"`
	WBSElementExternalID           string `json:"WBSElementExternalID"`
	WBSElementShortID              string `json:"WBSElementShortID"`
	WBSDescription                 string `json:"WBSDescription"`
	WBSElementLangBsdDescription   string `json:"WBSElementLangBsdDescription"`
	ResponsiblePerson              string `json:"ResponsiblePerson"`
	ResponsiblePersonName          string `json:"ResponsiblePersonName"`
	ApplicantCode                  string `json:"ApplicantCode"`
	ApplicantName                  string `json:"ApplicantName"`
	CompanyCode                    string `json:"CompanyCode"`
	BusinessArea                   string `json:"BusinessArea"`
	ControllingArea                string `json:"ControllingArea"`
	FunctionalArea                 string `json:"FunctionalArea"`
	ProfitCenter                   string `json:"ProfitCenter"`
	ResponsibleCostCenter          string `json:"ResponsibleCostCenter"`
	Plant                          string `json:"Plant"`
	FreeDefinedTableFieldSemantic  string `json:"FreeDefinedTableFieldSemantic"`
	FactoryCalendar                string `json:"FactoryCalendar"`
	PriorityCode                   string `json:"PriorityCode"`
	Currency                       string `json:"Currency"`
	CostingSheet                   string `json:"CostingSheet"`
	CostCenter                     string `json:"CostCenter"`
	RequestingCostCenter           string `json:"RequestingCostCenter"`
	ProjectInternalID              string `json:"ProjectInternalID"`
	WBSElementIsBillingElement     bool   `json:"WBSElementIsBillingElement"`
	InvestmentProfile              string `json:"InvestmentProfile"`
	WBSIsStatisticalWBSElement     bool   `json:"WBSIsStatisticalWBSElement"`
	WBSIsAccountAssignmentElement  bool   `json:"WBSIsAccountAssignmentElement"`
	ProjectType                    string `json:"ProjectType"`
	WBSElementIsPlanningElement    bool   `json:"WBSElementIsPlanningElement"`
	WorkCenterLocation             string `json:"WorkCenterLocation"`
	ResultAnalysisInternalID       string `json:"ResultAnalysisInternalID"`
	TaxJurisdiction                string `json:"TaxJurisdiction"`
	FunctionalLocation             string `json:"FunctionalLocation"`
	CreationDate                   string `json:"CreationDate"`
	LastChangeDate                 string `json:"LastChangeDate"`
	WBSIsMarkedForIntegratedPlng   bool   `json:"WBSIsMarkedForIntegratedPlng"`
	Equipment                      string `json:"Equipment"`
	ProjectObjectChangeNumber      string `json:"ProjectObjectChangeNumber"`
	WBSElementHierarchyLevel       int    `json:"WBSElementHierarchyLevel"`
	ReferenceElement               string `json:"ReferenceElement"`
	ProjInterestCalcProfile        string `json:"ProjInterestCalcProfile"`
	Language                       string `json:"Language"`
	IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
	BasicStartDate                 string `json:"BasicStartDate"`
	ForecastedStartDate            string `json:"ForecastedStartDate"`
	ActualStartDate                string `json:"ActualStartDate"`
	BasicEndDate                   string `json:"BasicEndDate"`
	ForecastedEndDate              string `json:"ForecastedEndDate"`
	ActualEndDate                  string `json:"ActualEndDate"`
	BasicDuration                  string `json:"BasicDuration"`
	BasicDurationUnit              string `json:"BasicDurationUnit"`
	ForecastedDuration             string `json:"ForecastedDuration"`
	ForecastedDurationUnit         string `json:"ForecastedDurationUnit"`
	ActualDuration                 string `json:"ActualDuration"`
	ActualDurationUnit             string `json:"ActualDurationUnit"`
	SchedldBasicEarliestStartDate  string `json:"SchedldBasicEarliestStartDate"`
	ScheduledBasicLatestEndDate    string `json:"ScheduledBasicLatestEndDate"`
	SchedldFcstdEarliestStartDate  string `json:"SchedldFcstdEarliestStartDate"`
	LatestSchedldFcstdEndDate      string `json:"LatestSchedldFcstdEndDate"`
	TentativeActualStartDate       string `json:"TentativeActualStartDate"`
	TentativeActualEndDate         string `json:"TentativeActualEndDate"`
	JointVenture                   string `json:"JointVenture"`
	JointVentureEquityType         string `json:"JointVentureEquityType"`
	JntVntrProjectType             string `json:"JntVntrProjectType"`
	SubProject                     string `json:"SubProject"`
	InvestmentReason               string `json:"InvestmentReason"`
	InvestmentScale                string `json:"InvestmentScale"`
	EnvironmentalInvestmentReason  string `json:"EnvironmentalInvestmentReason"`
	RequestingCompanyCode          string `json:"RequestingCompanyCode"`
	NetworkAssignmentType          string `json:"NetworkAssignmentType"`
	CostObject                     string `json:"CostObject"`
	WBSElementParentInternalID     string `json:"WBSElementParentInternalID"`
	WBSElementChildInternalID      string `json:"WBSElementChildInternalID"`
	LeftSiblingWBSElmntInternalID  string `json:"LeftSiblingWBSElmntInternalID"`
	RightSiblingWBSElmntInternalID string `json:"RightSiblingWBSElmntInternalID"`
	FreeDefinedAttribute01         string `json:"FreeDefinedAttribute01"`
	FreeDefinedAttribute02         string `json:"FreeDefinedAttribute02"`
	FreeDefinedAttribute03         string `json:"FreeDefinedAttribute03"`
	FreeDefinedAttribute04         string `json:"FreeDefinedAttribute04"`
	FreeDefinedQuantity1           string `json:"FreeDefinedQuantity1"`
	FreeDefinedQuantity1Unit       string `json:"FreeDefinedQuantity1Unit"`
	FreeDefinedQuantity2           string `json:"FreeDefinedQuantity2"`
	FreeDefinedQuantity2Unit       string `json:"FreeDefinedQuantity2Unit"`
	FreeDefinedAmount1             string `json:"FreeDefinedAmount1"`
	FreeDefinedAmount1Currency     string `json:"FreeDefinedAmount1Currency"`
	FreeDefinedAmount2             string `json:"FreeDefinedAmount2"`
	FreeDefinedAmount2Currency     string `json:"FreeDefinedAmount2Currency"`
	FreeDefinedDate1               string `json:"FreeDefinedDate1"`
	FreeDefinedDate2               string `json:"FreeDefinedDate2"`
	FreeDefinedIndicator1          bool   `json:"FreeDefinedIndicator1"`
	FreeDefinedIndicator2          bool   `json:"FreeDefinedIndicator2"`
}