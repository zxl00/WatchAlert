package services

import (
	service2 "watchAlert/pkg/community/aws/cloudwatch/service"
	"watchAlert/pkg/community/aws/service"
	"watchAlert/pkg/ctx"
)

var (
	DatasourceService       InterDatasourceService
	AuditLogService         InterAuditLogService
	DashboardService        InterDashboardService
	DutyManageService       InterDutyManageService
	DutyCalendarService     InterDutyCalendarService
	EventService            InterEventService
	NoticeService           InterNoticeService
	NoticeTmplService       InterNoticeTmplService
	RuleService             InterRuleService
	RuleGroupService        InterRuleGroupService
	RuleTmplService         InterRuleTmplService
	SilenceService          InterSilenceService
	TenantService           InterTenantService
	UserService             InterUserService
	LdapService             InterLdapService
	UserRoleService         InterUserRoleService
	AlertService            InterAlertService
	RuleTmplGroupService    InterRuleTmplGroupService
	UserPermissionService   InterUserPermissionService
	AWSRegionService        service.InterAwsRegionService
	AWSCloudWatchService    service2.InterAwsCloudWatchService
	AWSCloudWatchRdsService service2.InterAwsRdsService
	SettingService          InterSettingService
	ClientService           InterClientService
	MonitorService          InterMonitorService
)

func NewServices(ctx *ctx.Context) {
	DatasourceService = newInterDatasourceService(ctx)
	AuditLogService = newInterAuditLogService(ctx)
	DashboardService = newInterDashboardService(ctx)
	DutyManageService = newInterDutyManageService(ctx)
	DutyCalendarService = newInterDutyCalendarService(ctx)
	EventService = newInterEventService(ctx)
	NoticeService = newInterAlertNoticeService(ctx)
	NoticeTmplService = newInterNoticeTmplService(ctx)
	RuleService = newInterRuleService(ctx)
	RuleGroupService = newInterRuleGroupService(ctx)
	RuleTmplService = newInterRuleTmplService(ctx)
	RuleTmplGroupService = newInterRuleTmplGroupService(ctx)
	SilenceService = newInterSilenceService(ctx)
	TenantService = newInterTenantService(ctx)
	UserService = newInterUserService(ctx)
	LdapService = newInterLdapService(ctx)
	UserRoleService = newInterUserRoleService(ctx)
	AlertService = newInterAlertService(ctx)
	UserPermissionService = newInterUserPermissionService(ctx)
	AWSRegionService = service.NewInterAwsRegionService(ctx)
	AWSCloudWatchService = service2.NewInterAwsCloudWatchService(ctx)
	AWSCloudWatchRdsService = service2.NewInterAWSRdsService(ctx)
	SettingService = newInterSettingService(ctx)
	ClientService = newInterClientService(ctx)
	MonitorService = newInterMonitorService(ctx)
}
