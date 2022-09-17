package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	SalesOrder    struct {
		SalesOrder     string `json:"document_no"`
		Plant          string `json:"deliver_to"`
		OrderQuantity  string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		NetPriceAmount string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	ProductionOrder struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                string   `json:"api_schema"`
	Accepter                 []string `json:"accepter"`
	MaterialCode             string   `json:"material_code"`
	Supplier                 string   `json:"plant/supplier"`
	Stock                    string   `json:"stock"`
	SalesOrderType           string   `json:"document_type"`
	SONumber                 string   `json:"document_no"`
	ScheduleLineDeliveryDate string   `json:"planned_date"`
	ValidatedDate            string   `json:"validated_date"`
	Deleted                  bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey string `json:"connection_key"`
	Reult         bool   `json:"reult"`
	RediKey       string `json:"redi_key"`
	Filepath      string `json:"filepath"`
	Project       struct {
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
		TaxJurisdiction               string `json:"TaxJurisdiction"`
		ResponsiblePerson             string `json:"ResponsiblePerson"`
		ResponsiblePersonName         string `json:"ResponsiblePersonName"`
		ApplicantCode                 string `json:"ApplicantCode"`
		ApplicantName                 string `json:"ApplicantName"`
		CreatedByUser                 string `json:"CreatedByUser"`
		CreationDate                  string `json:"CreationDate"`
		LastChangedByUser             string `json:"LastChangedByUser"`
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
		ControllingObjectClass        string `json:"ControllingObjectClass"`
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
		DibutionChannel               string `json:"DibutionChannel"`
		Language                      string `json:"Language"`
		WBSElementExternalID          string `json:"WBSElementExternalID"`
		Division                      string `json:"Division"`
		DynItemProcessorPrfl          string `json:"DynItemProcessorPrfl"`
		JointVenture                  string `json:"JointVenture"`
		JointVentureCostRecoveryCode  string `json:"JointVentureCostRecoveryCode"`
		JointVentureEquityType        string `json:"JointVentureEquityType"`
		JointVentureObjectType        string `json:"JointVentureObjectType"`
		JntIntrstBillgClass           string `json:"JntIntrstBillgClass"`
		JntIntrstBillgSubClass        string `json:"JntIntrstBillgSubClass"`
		StatusProfile                 string `json:"StatusProfile"`
		WBSStatusProfile              string `json:"WBSStatusProfile"`
		SimulationProfile             string `json:"SimulationProfile"`
		SchedulingScenario            string `json:"SchedulingScenario"`
		DibutionProfile               string `json:"DibutionProfile"`
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
		WBSElement                    struct {
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
			CreatedByUser                  string `json:"CreatedByUser"`
			CreationDate                   string `json:"CreationDate"`
			LastChangedByUser              string `json:"LastChangedByUser"`
			LastChangeDate                 string `json:"LastChangeDate"`
			RespCostCenterControllingArea  string `json:"RespCostCenterControllingArea"`
			WBSIsMarkedForIntegratedPlng   bool   `json:"WBSIsMarkedForIntegratedPlng"`
			Equipment                      string `json:"Equipment"`
			ProjectObjectChangeNumber      string `json:"ProjectObjectChangeNumber"`
			WBSElementHierarchyLevel       int    `json:"WBSElementHierarchyLevel"`
			OverheadCode                   string `json:"OverheadCode"`
			ReferenceElement               string `json:"ReferenceElement"`
			ProjInterestCalcProfile        string `json:"ProjInterestCalcProfile"`
			ProgressAnlysAggregationWeight string `json:"ProgressAnlysAggregationWeight"`
			ReqgCostCenterControllingArea  string `json:"ReqgCostCenterControllingArea"`
			Language                       string `json:"Language"`
			IsMarkedForDeletion            bool   `json:"IsMarkedForDeletion"`
			WBSElementIsGroupingWBSElement string `json:"WBSElementIsGroupingWBSElement"`
			WBSElementIsUsedInProjSmmry    bool   `json:"WBSElementIsUsedInProjSmmry"`
			CostingVariant                 string `json:"CostingVariant"`
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
			JointVentureCostRecoveryCode   string `json:"JointVentureCostRecoveryCode"`
			JointVentureEquityType         string `json:"JointVentureEquityType"`
			JntVntrProjectType             string `json:"JntVntrProjectType"`
			JntIntrstBillgClass            string `json:"JntIntrstBillgClass"`
			JntIntrstBillgSubClass         string `json:"JntIntrstBillgSubClass"`
			SubProject                     string `json:"SubProject"`
			InvestmentReason               string `json:"InvestmentReason"`
			InvestmentScale                string `json:"InvestmentScale"`
			EnvironmentalInvestmentReason  string `json:"EnvironmentalInvestmentReason"`
			RequestingCompanyCode          string `json:"RequestingCompanyCode"`
			NetworkAssignmentType          string `json:"NetworkAssignmentType"`
			CostObject                     string `json:"CostObject"`
			BillingPlan                    string `json:"BillingPlan"`
			ControllingObjectClass         string `json:"ControllingObjectClass"`
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
		} `json:"WBSElement"`
	} `json:"Project"`
	APISchema string   `json:"api_schema"`
	Accepter  []string `json:"accepter"`
	ProjectID string   `json:"project_id"`
	Deleted   bool     `json:"deleted"`
}
