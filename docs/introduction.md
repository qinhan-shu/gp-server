# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/analysis.proto](#proto/analysis.proto)
    - [AnalysisByDifficultyReq](#protocol.AnalysisByDifficultyReq)
    - [AnalysisByDifficultyResp](#protocol.AnalysisByDifficultyResp)
    - [AnalysisByDifficultyResp.LineEntry](#protocol.AnalysisByDifficultyResp.LineEntry)
    - [AnalysisByDifficultyResp.PieEntry](#protocol.AnalysisByDifficultyResp.PieEntry)
    - [AnalysisByTagsReq](#protocol.AnalysisByTagsReq)
    - [AnalysisByTagsResp](#protocol.AnalysisByTagsResp)
    - [AnalysisByTagsResp.LineEntry](#protocol.AnalysisByTagsResp.LineEntry)
    - [AnalysisByTagsResp.PieEntry](#protocol.AnalysisByTagsResp.PieEntry)
  
  
  
  

- [proto/announcement.proto](#proto/announcement.proto)
    - [AddAnnouncementReq](#protocol.AddAnnouncementReq)
    - [AddAnnouncementResp](#protocol.AddAnnouncementResp)
    - [AnnouncementDetailReq](#protocol.AnnouncementDetailReq)
    - [AnnouncementDetailResp](#protocol.AnnouncementDetailResp)
    - [AnnouncementsReq](#protocol.AnnouncementsReq)
    - [AnnouncementsResp](#protocol.AnnouncementsResp)
    - [DelAnnouncementReq](#protocol.DelAnnouncementReq)
    - [DelAnnouncementResp](#protocol.DelAnnouncementResp)
    - [EditAnnouncementReq](#protocol.EditAnnouncementReq)
    - [EditAnnouncementResp](#protocol.EditAnnouncementResp)
  
  
  
  

- [proto/class_enter.proto](#proto/class_enter.proto)
    - [ApplyEnterRequestReq](#protocol.ApplyEnterRequestReq)
    - [ApplyEnterRequestResp](#protocol.ApplyEnterRequestResp)
    - [EnterClassReq](#protocol.EnterClassReq)
    - [EnterClassResp](#protocol.EnterClassResp)
    - [GetClasserByUserID](#protocol.GetClasserByUserID)
    - [GetMemberReq](#protocol.GetMemberReq)
    - [GetMemberResp](#protocol.GetMemberResp)
    - [QuitClassReq](#protocol.QuitClassReq)
    - [QuitClassResp](#protocol.QuitClassResp)
    - [SearchClassesReq](#protocol.SearchClassesReq)
    - [SearchClassesResp](#protocol.SearchClassesResp)
  
  
  
  

- [proto/class_manage.proto](#proto/class_manage.proto)
    - [AddClassReq](#protocol.AddClassReq)
    - [AddClassResp](#protocol.AddClassResp)
    - [EditClassReq](#protocol.EditClassReq)
    - [EditClassResp](#protocol.EditClassResp)
    - [GetClassByIDReq](#protocol.GetClassByIDReq)
    - [GetClassByIDResp](#protocol.GetClassByIDResp)
    - [GetClassesReq](#protocol.GetClassesReq)
    - [GetClassesResp](#protocol.GetClassesResp)
    - [MemberManageReq](#protocol.MemberManageReq)
    - [MemberManageResp](#protocol.MemberManageResp)
  
    - [MemberManageReq.ManageType](#protocol.MemberManageReq.ManageType)
  
  
  

- [proto/common.proto](#proto/common.proto)
    - [Announcement](#protocol.Announcement)
    - [Class](#protocol.Class)
    - [ClassMember](#protocol.ClassMember)
    - [Match](#protocol.Match)
    - [Paper](#protocol.Paper)
    - [Problem](#protocol.Problem)
    - [ProblemExample](#protocol.ProblemExample)
    - [RankItem](#protocol.RankItem)
    - [SubmitRecord](#protocol.SubmitRecord)
    - [UserInfo](#protocol.UserInfo)
  
  
  
  

- [proto/conf.proto](#proto/conf.proto)
    - [Config](#protocol.Config)
    - [Config.DifficultyEntry](#protocol.Config.DifficultyEntry)
    - [Config.TagsEntry](#protocol.Config.TagsEntry)
    - [JudgeLanguage](#protocol.JudgeLanguage)
    - [JudgeLanguage.LanguageEntry](#protocol.JudgeLanguage.LanguageEntry)
    - [JudgeResults](#protocol.JudgeResults)
    - [JudgeResults.JudgeResultsEntry](#protocol.JudgeResults.JudgeResultsEntry)
    - [PaperComposeAlgorithm](#protocol.PaperComposeAlgorithm)
    - [PaperComposeAlgorithm.AlgorithmEntry](#protocol.PaperComposeAlgorithm.AlgorithmEntry)
    - [UserRole](#protocol.UserRole)
    - [UserRole.RoleEntry](#protocol.UserRole.RoleEntry)
  
  
  
  

- [proto/file.proto](#proto/file.proto)
    - [File](#protocol.File)
  
  
  
  

- [proto/judge.proto](#proto/judge.proto)
    - [JudgeRequest](#protocol.JudgeRequest)
    - [JudgeResponse](#protocol.JudgeResponse)
    - [JudgeResult](#protocol.JudgeResult)
  
    - [JudgeResult.Result](#protocol.JudgeResult.Result)
  
  
  

- [proto/login.proto](#proto/login.proto)
    - [LoginReq](#protocol.LoginReq)
    - [LoginResp](#protocol.LoginResp)
    - [LogoutResp](#protocol.LogoutResp)
    - [RegisterReq](#protocol.RegisterReq)
    - [RegisterResp](#protocol.RegisterResp)
  
  
  
  

- [proto/match.proto](#proto/match.proto)
    - [EditMatchReq](#protocol.EditMatchReq)
    - [EditMatchResp](#protocol.EditMatchResp)
    - [GetMatchByIDReq](#protocol.GetMatchByIDReq)
    - [GetMatchByIDResp](#protocol.GetMatchByIDResp)
    - [GetMatchesReq](#protocol.GetMatchesReq)
    - [GetMatchesResp](#protocol.GetMatchesResp)
    - [GetPaperByIDReq](#protocol.GetPaperByIDReq)
    - [GetPaperByIDResp](#protocol.GetPaperByIDResp)
    - [NewMatchReq](#protocol.NewMatchReq)
    - [NewMatchResp](#protocol.NewMatchResp)
  
  
  
  

- [proto/paper.proto](#proto/paper.proto)
    - [ManualModifyPaperReq](#protocol.ManualModifyPaperReq)
    - [ManualModifyPaperResp](#protocol.ManualModifyPaperResp)
    - [NewPaperReq](#protocol.NewPaperReq)
    - [NewPaperResp](#protocol.NewPaperResp)
  
    - [ManualModifyPaperReq.ModifyType](#protocol.ManualModifyPaperReq.ModifyType)
  
  
  

- [proto/problem_manage.proto](#proto/problem_manage.proto)
    - [AddProblemReq](#protocol.AddProblemReq)
    - [AddProblemResp](#protocol.AddProblemResp)
    - [EditProblemReq](#protocol.EditProblemReq)
    - [EditProblemResp](#protocol.EditProblemResp)
    - [GetProblemByIDReq](#protocol.GetProblemByIDReq)
    - [GetProblemByIDResp](#protocol.GetProblemByIDResp)
    - [GetProblemsReq](#protocol.GetProblemsReq)
    - [GetProblemsResp](#protocol.GetProblemsResp)
  
  
  
  

- [proto/rank.proto](#proto/rank.proto)
    - [RankListReq](#protocol.RankListReq)
    - [RankListResp](#protocol.RankListResp)
  
  
  
  

- [proto/status.proto](#proto/status.proto)
    - [Status](#protocol.Status)
  
    - [Code](#protocol.Code)
  
  
  

- [proto/tag.proto](#proto/tag.proto)
    - [AddTagReq](#protocol.AddTagReq)
    - [AddTagResp](#protocol.AddTagResp)
    - [GetTagsReq](#protocol.GetTagsReq)
    - [GetTagsResp](#protocol.GetTagsResp)
    - [GetTagsResp.TagsEntry](#protocol.GetTagsResp.TagsEntry)
    - [UpdateTagReq](#protocol.UpdateTagReq)
    - [UpdateTagResp](#protocol.UpdateTagResp)
  
  
  
  

- [proto/user_manage.proto](#proto/user_manage.proto)
    - [AddUsersReq](#protocol.AddUsersReq)
    - [AddUsersResp](#protocol.AddUsersResp)
    - [DelUsersReq](#protocol.DelUsersReq)
    - [DelUsersResp](#protocol.DelUsersResp)
    - [GetSubmitRecordReq](#protocol.GetSubmitRecordReq)
    - [GetSubmitRecordResp](#protocol.GetSubmitRecordResp)
    - [GetUsersReq](#protocol.GetUsersReq)
    - [GetUsersResp](#protocol.GetUsersResp)
    - [UpdateUsersReq](#protocol.UpdateUsersReq)
    - [UpdateUsersResp](#protocol.UpdateUsersResp)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/analysis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/analysis.proto



<a name="protocol.AnalysisByDifficultyReq"></a>

### AnalysisByDifficultyReq
根据难度进行数据分析


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| startTime | [int64](#int64) |  |  |
| endTime | [int64](#int64) |  |  |






<a name="protocol.AnalysisByDifficultyResp"></a>

### AnalysisByDifficultyResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| line | [AnalysisByDifficultyResp.LineEntry](#protocol.AnalysisByDifficultyResp.LineEntry) | repeated |  |
| pie | [AnalysisByDifficultyResp.PieEntry](#protocol.AnalysisByDifficultyResp.PieEntry) | repeated |  |






<a name="protocol.AnalysisByDifficultyResp.LineEntry"></a>

### AnalysisByDifficultyResp.LineEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [double](#double) |  |  |






<a name="protocol.AnalysisByDifficultyResp.PieEntry"></a>

### AnalysisByDifficultyResp.PieEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int64](#int64) |  |  |






<a name="protocol.AnalysisByTagsReq"></a>

### AnalysisByTagsReq
根据tags进行数据分析


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| startTime | [int64](#int64) |  |  |
| endTime | [int64](#int64) |  |  |
| tags | [int64](#int64) | repeated |  |






<a name="protocol.AnalysisByTagsResp"></a>

### AnalysisByTagsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| line | [AnalysisByTagsResp.LineEntry](#protocol.AnalysisByTagsResp.LineEntry) | repeated |  |
| pie | [AnalysisByTagsResp.PieEntry](#protocol.AnalysisByTagsResp.PieEntry) | repeated |  |






<a name="protocol.AnalysisByTagsResp.LineEntry"></a>

### AnalysisByTagsResp.LineEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [double](#double) |  |  |






<a name="protocol.AnalysisByTagsResp.PieEntry"></a>

### AnalysisByTagsResp.PieEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/announcement.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/announcement.proto



<a name="protocol.AddAnnouncementReq"></a>

### AddAnnouncementReq
增加新公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| announcement | [Announcement](#protocol.Announcement) |  |  |






<a name="protocol.AddAnnouncementResp"></a>

### AddAnnouncementResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.AnnouncementDetailReq"></a>

### AnnouncementDetailReq
获取公告具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.AnnouncementDetailResp"></a>

### AnnouncementDetailResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| announcement | [Announcement](#protocol.Announcement) |  |  |






<a name="protocol.AnnouncementsReq"></a>

### AnnouncementsReq
获取所有公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.AnnouncementsResp"></a>

### AnnouncementsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| announcements | [Announcement](#protocol.Announcement) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.DelAnnouncementReq"></a>

### DelAnnouncementReq
删除公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.DelAnnouncementResp"></a>

### DelAnnouncementResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditAnnouncementReq"></a>

### EditAnnouncementReq
修改公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| announcement | [Announcement](#protocol.Announcement) |  |  |






<a name="protocol.EditAnnouncementResp"></a>

### EditAnnouncementResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 

 

 

 



<a name="proto/class_enter.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/class_enter.proto



<a name="protocol.ApplyEnterRequestReq"></a>

### ApplyEnterRequestReq
教师批准/拒绝学生进入班级的申请


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| is_apply | [bool](#bool) |  | 是否同意 |






<a name="protocol.ApplyEnterRequestResp"></a>

### ApplyEnterRequestResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EnterClassReq"></a>

### EnterClassReq
申请进入班级


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [int64](#int64) |  |  |






<a name="protocol.EnterClassResp"></a>

### EnterClassResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetClasserByUserID"></a>

### GetClasserByUserID
通过学生id获取班级


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| classes | [Class](#protocol.Class) | repeated |  |






<a name="protocol.GetMemberReq"></a>

### GetMemberReq
获取班级成员
包括了查看未批准的进入班级的请求
当学生申请进入班级的时候，就会在user——class表中插入一条记录，对于需要进行教师认证的班级，其is-check字段为false
这个请求返回user-class表中所有的记录，包括通过的和未通过的，is-check字段在ClassMember中


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [int64](#int64) |  |  |






<a name="protocol.GetMemberResp"></a>

### GetMemberResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| members | [ClassMember](#protocol.ClassMember) | repeated |  |






<a name="protocol.QuitClassReq"></a>

### QuitClassReq
申请退出班级


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | [int64](#int64) |  |  |






<a name="protocol.QuitClassResp"></a>

### QuitClassResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.SearchClassesReq"></a>

### SearchClassesReq
搜索班级


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| keyword | [string](#string) |  |  |






<a name="protocol.SearchClassesResp"></a>

### SearchClassesResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| classes | [Class](#protocol.Class) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/class_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/class_manage.proto



<a name="protocol.AddClassReq"></a>

### AddClassReq
新增班级


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [Class](#protocol.Class) |  |  |






<a name="protocol.AddClassResp"></a>

### AddClassResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditClassReq"></a>

### EditClassReq
修改班级信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class | [Class](#protocol.Class) |  | 只需要填充id以及改变的field |






<a name="protocol.EditClassResp"></a>

### EditClassResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetClassByIDReq"></a>

### GetClassByIDReq
根据ID获得班级具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetClassByIDResp"></a>

### GetClassByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| class | [Class](#protocol.Class) |  |  |






<a name="protocol.GetClassesReq"></a>

### GetClassesReq
批量获取班级信息(只包括基础的信息，班级名称)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetClassesResp"></a>

### GetClassesResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| classes | [Class](#protocol.Class) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.MemberManageReq"></a>

### MemberManageReq
小组成员管理


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| manage_type | [MemberManageReq.ManageType](#protocol.MemberManageReq.ManageType) |  |  |
| class_id | [int64](#int64) |  |  |
| member_id | [int64](#int64) |  |  |






<a name="protocol.MemberManageResp"></a>

### MemberManageResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 


<a name="protocol.MemberManageReq.ManageType"></a>

### MemberManageReq.ManageType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DELETE | 0 | 删除小组成员 |
| SET_ADMINISTRATOR | 1 | 设置成管理员 |
| CANCEL_ADMINISTRATOR | 2 | 取消管理员 |


 

 

 



<a name="proto/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/common.proto



<a name="protocol.Announcement"></a>

### Announcement
Announcement : 公告，包括班级公告和全局公告


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| publisher | [string](#string) |  | 发布人的姓名 |
| title | [string](#string) |  | 公告标题 |
| detail | [string](#string) |  |  |
| create_time | [int64](#int64) |  |  |
| last_update_time | [int64](#int64) |  |  |






<a name="protocol.Class"></a>

### Class
Class : 班级信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 班级id |
| tutor | [string](#string) |  | 导师的姓名 |
| name | [string](#string) |  | 班级名称 |
| introduction | [string](#string) |  | 班级简介 |
| number | [int64](#int64) |  | 班级人数 |
| is_check | [bool](#bool) |  | 加入班级设置：false（无需审核，运行任何人进入），true（需要导师审核） |
| create_time | [int64](#int64) |  | 班级创建时间 |
| announcements | [Announcement](#protocol.Announcement) | repeated | 班级公告 |






<a name="protocol.ClassMember"></a>

### ClassMember



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| is_administrator | [bool](#bool) |  |  |
| is_checked | [bool](#bool) |  |  |






<a name="protocol.Match"></a>

### Match
Match : 比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| is_public | [bool](#bool) |  | 是否是公开比赛 |
| start_time | [int64](#int64) |  | 开始时间 |
| end_time | [int64](#int64) |  | 结束长度 |
| is_over | [bool](#bool) |  | 是否结束 |
| name | [string](#string) |  | 比赛名称 |
| introduction | [string](#string) |  | 比赛简介 |
| paper_id | [int64](#int64) |  | 试卷id |






<a name="protocol.Paper"></a>

### Paper
Paper : 试卷


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| problems | [Problem](#protocol.Problem) | repeated | 题目 |
| difficulty | [int64](#int64) |  | 组卷的参数 |
| problem_num | [int64](#int64) |  |  |
| tags | [int64](#int64) | repeated |  |






<a name="protocol.Problem"></a>

### Problem
Problem : 题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 题目id |
| title | [string](#string) |  | 题目标题 |
| description | [string](#string) |  | 题目描述 |
| in | [string](#string) |  | 输入 |
| out | [string](#string) |  | 输出 |
| hint | [string](#string) |  | 题目提示 |
| in_out_examples | [ProblemExample](#protocol.ProblemExample) | repeated | 输入输出样例 |
| judge_limit_time | [int64](#int64) |  | 判题时间限制 |
| judge_limit_mem | [int64](#int64) |  | 判题内存限制 |
| tags | [int64](#int64) | repeated | 题目标签 |
| difficulty | [int64](#int64) |  | 难度 |
| cognition | [int64](#int64) |  | 认知程度 |
| submit_time | [int64](#int64) |  | 提交次数 |
| accept_time | [int64](#int64) |  | 通过次数 |
| create_time | [int64](#int64) |  | 创建时间 |
| publisher | [string](#string) |  | 发布人 |
| judge_file | [string](#string) |  | 判题文件 |






<a name="protocol.ProblemExample"></a>

### ProblemExample
ProblemExample : 题目输入输出样例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| index | [int64](#int64) |  |  |
| input | [string](#string) |  |  |
| output | [string](#string) |  |  |






<a name="protocol.RankItem"></a>

### RankItem
RankItem : 排名item


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ranking | [int64](#int64) |  | 排名 |
| user_id | [int64](#int64) |  | 用户id |
| name | [string](#string) |  | 用户姓名 |
| pass_num | [int64](#int64) |  | 通过题目数量 |
| submit_num | [int64](#int64) |  | 提交次数 |






<a name="protocol.SubmitRecord"></a>

### SubmitRecord
SubmitRecord : 提交情况


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem_id | [int64](#int64) |  | 题目 |
| user_id | [int64](#int64) |  | 用户id |
| submit_time | [int64](#int64) |  | 提交时间戳 |
| is_pass | [bool](#bool) |  | 是否通过 |
| running_time | [int64](#int64) |  |  |
| running_mem | [int64](#int64) |  |  |
| code | [string](#string) |  |  |
| running_lan | [int64](#int64) |  |  |






<a name="protocol.UserInfo"></a>

### UserInfo
UserInfo : 用户基本信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| role | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| sex | [bool](#bool) |  |  |
| phone | [string](#string) |  |  |
| email | [string](#string) |  |  |
| school | [string](#string) |  |  |
| last_login | [int64](#int64) |  |  |
| create | [int64](#int64) |  |  |
| account | [string](#string) |  | 这两个字段只有在用户管理中的新增用户才会用到, 客户端向服务端发送数据是填充 |
| password | [string](#string) |  |  |





 

 

 

 



<a name="proto/conf.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/conf.proto



<a name="protocol.Config"></a>

### Config



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| tags | [Config.TagsEntry](#protocol.Config.TagsEntry) | repeated |  |
| difficulty | [Config.DifficultyEntry](#protocol.Config.DifficultyEntry) | repeated |  |






<a name="protocol.Config.DifficultyEntry"></a>

### Config.DifficultyEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="protocol.Config.TagsEntry"></a>

### Config.TagsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="protocol.JudgeLanguage"></a>

### JudgeLanguage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| language | [JudgeLanguage.LanguageEntry](#protocol.JudgeLanguage.LanguageEntry) | repeated |  |






<a name="protocol.JudgeLanguage.LanguageEntry"></a>

### JudgeLanguage.LanguageEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="protocol.JudgeResults"></a>

### JudgeResults



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| judge_results | [JudgeResults.JudgeResultsEntry](#protocol.JudgeResults.JudgeResultsEntry) | repeated |  |






<a name="protocol.JudgeResults.JudgeResultsEntry"></a>

### JudgeResults.JudgeResultsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="protocol.PaperComposeAlgorithm"></a>

### PaperComposeAlgorithm
PaperCompose 组卷算法


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| algorithm | [PaperComposeAlgorithm.AlgorithmEntry](#protocol.PaperComposeAlgorithm.AlgorithmEntry) | repeated |  |






<a name="protocol.PaperComposeAlgorithm.AlgorithmEntry"></a>

### PaperComposeAlgorithm.AlgorithmEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="protocol.UserRole"></a>

### UserRole



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| role | [UserRole.RoleEntry](#protocol.UserRole.RoleEntry) | repeated |  |






<a name="protocol.UserRole.RoleEntry"></a>

### UserRole.RoleEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |





 

 

 

 



<a name="proto/file.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/file.proto



<a name="protocol.File"></a>

### File



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| file_id | [string](#string) |  |  |





 

 

 

 



<a name="proto/judge.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/judge.proto



<a name="protocol.JudgeRequest"></a>

### JudgeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| src | [string](#string) |  |  |
| language | [int64](#int64) |  |  |






<a name="protocol.JudgeResponse"></a>

### JudgeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| result | [int64](#int64) |  |  |
| results | [JudgeResult](#protocol.JudgeResult) | repeated |  |






<a name="protocol.JudgeResult"></a>

### JudgeResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| judge_result | [JudgeResult.Result](#protocol.JudgeResult.Result) |  |  |
| cpu_time | [int64](#int64) |  | cpu time the process has used |
| real_time | [int64](#int64) |  | actual running time of the process |
| memory | [int64](#int64) |  | max vaule of memory used by the process |
| signal | [int64](#int64) |  | signal number |
| exit_code | [int64](#int64) |  | process&#39;s exit code |





 


<a name="protocol.JudgeResult.Result"></a>

### JudgeResult.Result
判题目结果

| Name | Number | Description |
| ---- | ------ | ----------- |
| SUCCESS | 0 | this means the answer is accepted |
| WRONG_ANSWER | -1 | this means the process exited normally, but the answer is wrong |
| CPU_TIME_LIMIT_EXCEEDED | 1 |  |
| REAL_TIME_LIMIT_EXCEEDED | 2 |  |
| MEMORY_LIMIT_EXCEEDED | 3 |  |
| RUNTIME_ERROR | 4 |  |
| SYSTEM_ERROR | 5 |  |


 

 

 



<a name="proto/login.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login.proto



<a name="protocol.LoginReq"></a>

### LoginReq
登陆 (post)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="protocol.LoginResp"></a>

### LoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| token | [string](#string) |  |  |
| user | [UserInfo](#protocol.UserInfo) |  | 用户信息 |






<a name="protocol.LogoutResp"></a>

### LogoutResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |






<a name="protocol.RegisterReq"></a>

### RegisterReq
登陆 (post)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [UserInfo](#protocol.UserInfo) |  |  |






<a name="protocol.RegisterResp"></a>

### RegisterResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 

 

 

 



<a name="proto/match.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/match.proto



<a name="protocol.EditMatchReq"></a>

### EditMatchReq
编辑比赛题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| match | [Match](#protocol.Match) |  |  |






<a name="protocol.EditMatchResp"></a>

### EditMatchResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetMatchByIDReq"></a>

### GetMatchByIDReq
拿到所有的比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetMatchByIDResp"></a>

### GetMatchByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| match | [Match](#protocol.Match) |  |  |






<a name="protocol.GetMatchesReq"></a>

### GetMatchesReq
拿到所有的比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetMatchesResp"></a>

### GetMatchesResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |
| matches | [Match](#protocol.Match) | repeated |  |






<a name="protocol.GetPaperByIDReq"></a>

### GetPaperByIDReq
获取比赛试卷


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetPaperByIDResp"></a>

### GetPaperByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| paper | [Paper](#protocol.Paper) |  |  |






<a name="protocol.NewMatchReq"></a>

### NewMatchReq
发布比赛


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paper_id | [int64](#int64) |  |  |
| match | [Match](#protocol.Match) |  |  |






<a name="protocol.NewMatchResp"></a>

### NewMatchResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 

 

 

 



<a name="proto/paper.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/paper.proto



<a name="protocol.ManualModifyPaperReq"></a>

### ManualModifyPaperReq
手动修改试卷


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| modifyType | [ManualModifyPaperReq.ModifyType](#protocol.ManualModifyPaperReq.ModifyType) |  |  |
| problem_id | [int64](#int64) |  |  |
| paper_id | [int64](#int64) |  |  |






<a name="protocol.ManualModifyPaperResp"></a>

### ManualModifyPaperResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.NewPaperReq"></a>

### NewPaperReq
创建试卷


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| paper | [Paper](#protocol.Paper) |  |  |
| algorithm | [int64](#int64) |  |  |






<a name="protocol.NewPaperResp"></a>

### NewPaperResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |
| paper | [Paper](#protocol.Paper) |  |  |





 


<a name="protocol.ManualModifyPaperReq.ModifyType"></a>

### ManualModifyPaperReq.ModifyType


| Name | Number | Description |
| ---- | ------ | ----------- |
| ADD | 0 |  |
| DEL | 1 |  |


 

 

 



<a name="proto/problem_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/problem_manage.proto



<a name="protocol.AddProblemReq"></a>

### AddProblemReq
新增题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  |  |






<a name="protocol.AddProblemResp"></a>

### AddProblemResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditProblemReq"></a>

### EditProblemReq
编辑题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  | 需要id |






<a name="protocol.EditProblemResp"></a>

### EditProblemResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetProblemByIDReq"></a>

### GetProblemByIDReq
根据ID获得题目具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetProblemByIDResp"></a>

### GetProblemByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| problem | [Problem](#protocol.Problem) |  |  |






<a name="protocol.GetProblemsReq"></a>

### GetProblemsReq
获取全部题目信息（只下发 id &amp; title &amp; difficulty &amp; pass_rate）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [int64](#int64) |  |  |
| keyword | [string](#string) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetProblemsResp"></a>

### GetProblemsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| problems | [Problem](#protocol.Problem) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/rank.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/rank.proto



<a name="protocol.RankListReq"></a>

### RankListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.RankListResp"></a>

### RankListResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| rank_items | [RankItem](#protocol.RankItem) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |
| pos | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/status.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/status.proto



<a name="protocol.Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [Code](#protocol.Code) |  |  |
| message | [string](#string) |  |  |





 


<a name="protocol.Code"></a>

### Code


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | ok |
| INTERNAL | 1 | 服务端内部错误 |
| DATA_LOSE | 2 | 数据序列化错误 |
| NO_TOKEN | 3 | 没有token |
| UNAUTHORIZATED | 4 | token错误 |
| PERMISSION_DENIED | 5 | 没有权限 |


 

 

 



<a name="proto/tag.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/tag.proto



<a name="protocol.AddTagReq"></a>

### AddTagReq
新增Tag


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [string](#string) |  |  |






<a name="protocol.AddTagResp"></a>

### AddTagResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetTagsReq"></a>

### GetTagsReq
获取所有tag


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetTagsResp"></a>

### GetTagsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| tags | [GetTagsResp.TagsEntry](#protocol.GetTagsResp.TagsEntry) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.GetTagsResp.TagsEntry"></a>

### GetTagsResp.TagsEntry



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [int64](#int64) |  |  |
| value | [string](#string) |  |  |






<a name="protocol.UpdateTagReq"></a>

### UpdateTagReq
修改tag


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| tag | [string](#string) |  |  |






<a name="protocol.UpdateTagResp"></a>

### UpdateTagResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |





 

 

 

 



<a name="proto/user_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/user_manage.proto



<a name="protocol.AddUsersReq"></a>

### AddUsersReq
批量增加用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserInfo](#protocol.UserInfo) | repeated |  |






<a name="protocol.AddUsersResp"></a>

### AddUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [UserInfo](#protocol.UserInfo) | repeated |  |
| fail | [UserInfo](#protocol.UserInfo) | repeated |  |






<a name="protocol.DelUsersReq"></a>

### DelUsersReq
批量删除用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users_id | [int64](#int64) | repeated |  |






<a name="protocol.DelUsersResp"></a>

### DelUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [int64](#int64) | repeated |  |
| fail | [int64](#int64) | repeated |  |






<a name="protocol.GetSubmitRecordReq"></a>

### GetSubmitRecordReq
获得做题记录


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| problem_id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |






<a name="protocol.GetSubmitRecordResp"></a>

### GetSubmitRecordResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| submit_records | [SubmitRecord](#protocol.SubmitRecord) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.GetUsersReq"></a>

### GetUsersReq
批量获取用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [int64](#int64) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetUsersResp"></a>

### GetUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| users | [UserInfo](#protocol.UserInfo) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |
| total | [int64](#int64) |  |  |






<a name="protocol.UpdateUsersReq"></a>

### UpdateUsersReq
批量修改用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserInfo](#protocol.UserInfo) | repeated | 只需要填充id以及改变的field |






<a name="protocol.UpdateUsersResp"></a>

### UpdateUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [UserInfo](#protocol.UserInfo) | repeated |  |
| fail | [UserInfo](#protocol.UserInfo) | repeated |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

