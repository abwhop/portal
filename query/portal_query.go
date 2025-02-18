package query

var NewsQuery = `{"query":"{news(limit: %d page: %d order_by: \"change_date\" order_sort:\"desc\") {id publish_date create_date change_date name log_id can_comment text create_date publish_date published rights img preview_text xml_id slider_file rubric {name id code} author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} views {count users {id name last_name second_name email position personal_number}} comments {id source_id parent_source_id text date_create author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width} vote_num repost_blog {id} tags{name} vote_num calendar_events{id title date_start date_end description location source_id entity_type created_by modified_by date_create date_update}}}"}`

var BlogPostQuery = `{"query":"{blog_posts (blog_post_type: \"all\" order_by:\"publish_date\" order_sort:\"desc\" limit: %d page: %d site_id: \"s1\") {id title vote_num blog_id text is_draft create_date publish_date img rights post_rights files {id link file_name original_name content_type size height width} author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} views {count users {id name last_name second_name email position personal_number}} comments {id source_id text date_create author {id name last_name second_name email position personal_number} likes {count users{id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width}} repost_blog {id title vote_num blog_id text is_draft create_date publish_date img rights post_rights files {id link file_name original_name content_type size height width} author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} views {count users {id name last_name second_name email position personal_number}} comments {id source_id text date_create author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width}}} repost_news {id name log_id can_comment text create_date publish_date published rights img preview_text xml_id slider_file rubric {name id code} author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} views {count users {id name last_name second_name email position personal_number}} comments {id source_id text date_create author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width}} repost_blog {id} views {count users {id name last_name second_name email position personal_number}} comments {id source_id text date_create author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width}}} files {id link file_name original_name content_type size height width} vote_num} vote {id title description author {id name last_name second_name email position personal_number} active date_from date_to questions {id sort question date_change active counter diagram required diagram_type question_type answers {id sort message field_type date_change active counter}} img {id link file_name original_name content_type size height width} date_change url vote_group {id name sort active hidden date_change title vote_single use_captcha site_id} views counter}}"}`

var BlogQuery = `{"query":"{blogs (limit: %d, page: %d, site_id: \"s1\") {id, name, description, date_create, author {id, name, last_name, second_name, email, position, personal_number, photo, login_ad}, subscribers {id, name, last_name, second_name, email, position, personal_number, photo, login_ad}}}","variables":{}}`

var AppointmentsQuery = `{"query": "{appointments(limit:%d page:%d order_by:\"id\" site_id: \"s1\" order_sort:\"desc\") {id name text create_date publish_date img preview_text likes {count users {id name last_name second_name email position personal_number}} views {count users {id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width} can_like log_id site_id can_comment rights author {id name last_name second_name email position personal_number} can_comment published publish_date create_date img preview_text rubric {id, name}, comments {id source_id parent_source_id text date_create author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}}}}"}`

var CommunityQuery = `{"query": "{workgroups (limit:%d page:%d site_id:\"s1\") {id, name, description, active, date_create, img, closed, visible, opened, project, user_is_member, group_is_favorite, project_date_start, project_date_end, subject {id, name, sort}, type {code, name, description}, author {id, login, active, name, last_name, second_name, email, gender, photo, company, department, position, birthday, company_id, personal_number}, files {id, link, file_name, original_name, content_type, size, height, width}, members {id, login, active, name, last_name, second_name, email, gender, photo, company, department, position, birthday, company_id, personal_number}, moderators {id, login, active, name, last_name, second_name, email, gender, photo, company, department, position, birthday, company_id, personal_number}, favorites {id, login, active, name, last_name, second_name, email, gender, photo, company, department, position, birthday, company_id, personal_number}, features {blog {write_post}}}}"}`

var CommunitySubjects = `{"query": "{workgroup_subjects {id, name, sort}}"}`

var CommunityTypes = `{"query": "{workgroup_types {id, name, sort}}"}`

//var UserQuery = `{"query":"{users (limit:1000, page:1) {id login active name last_name second_name email personal_mobile personal_phone gender photo company department position birthday company_id personal_number full_personal_number create_date update_date start_work_date factory_id education about_me hidden_fields manufactory_name department_name department_name_sp department_address chief_id login_ad favorites{workgroups{id name} news{id name} blog_posts{id title}} rights work_profile last_activity_date last_login rubrics{ appointment{id name} news{id name} vacancy{id name}} uf_site_id bx_department_id covid_qr_code covid_qr_code_decoded covid_qr_code_validation_data hidden_fields black_list_type black_list_message}}"}`

var UsersQuery = `{"query":"{users (limit:%d page:%d) {id login active name last_name second_name email personal_mobile personal_phone gender photo company department position birthday company_id personal_number full_personal_number create_date update_date start_work_date factory_id education about_me hidden_fields manufactory_name department_name department_name_sp department_address chief_id login_ad favorites { workgroups {id name } news {id name } blog_posts {id title } } rights work_profile last_activity_date last_login rubrics { appointment {id name } news {id name } vacancy {id name } } bx_department_id covid_qr_code covid_qr_code_decoded covid_qr_code_validation_data black_list_type black_list_message}}"}`

var FormQuery = `{"query": "{iblock (iblock_type: \"constructor_form\", order_sort: \"asc\", order_by: \"iblock_id\") {iblock_id, iblock_code, iblock_type, sort, name, active, properties {id, name, code, type, user_type, iblock_id,is_required, default_value, sort, active, multiple,xml_id, values {id, sort, name, xml_id, is_default, property_id}}, list_fields {iblock_id, field_id, sort, name, settings}}}"}`

var VotesQuery = `{"query": "{vote (limit:%d page:%d site_id: \"s1\") {id, title, description, author {id, name, last_name, second_name, email, position, personal_number}, active, date_from, date_to, questions {id, sort, question, date_change, active, counter, diagram, required, diagram_type, question_type, answers {id, sort, message, field_type, date_change, active, counter}}, img {id, link, file_name, original_name, content_type, size, height, width}, date_change, url, vote_group {id, name, sort, active, hidden, date_change, title, vote_single, use_captcha, site_id}, views, counter}}"}`

var VotesResultsQuery = `{"query": "{voteResults (limit:%d page:%d) {id, date, user {id, name, last_name, second_name, email, position, personal_number}, vote_id}}"}`

var UserSubscribedRubric = `{"query": "{users(id:%d) {rubrics{ news {id name code}}}}"}`

var UnSubscribeRubric = `{"query": "mutation{unsubscribeRubric (user_id:%d rubric_type:\"news\" rubric_id:%d) {success message}}"}`

var SubscribeRubric = `{"query": "mutation{subscribeRubric (user_id:%d rubric_type:\"news\" rubric_id:%d)  {success message}}"}`

var AddFavoriteQuery = `{"query": "mutation {addFavorites (user_id:%d news:[%d]) {news{id}}}"}`

var RemoveFavoriteQuery = `{"query": "mutation {removeFavorites (user_id:%d news:[%d]) {news{id}}}"}`

var SetCommentQuery = `{"query": "mutation {setComment (source_id:%d, type:\"news\" text:\"%s\", user_id:%d) {id source_id text date_create author {id name last_name second_name email position personal_number} likes {count users {id name last_name second_name email position personal_number}} files {id link file_name original_name content_type size height width}}}"}`

var SetLikeQuery = `{"query": "mutation {setLike(id:%d user_id:%d type:\"news\" cancelled:%t) {count users {id name last_name second_name email position personal_number}}}"}`

var SetViewQuery = `{"query": "mutation {setView(id:%d user_id:%d type:\"news\") {count users {id name last_name second_name email position personal_number}}}"}`

var SubscribeBlogQuery = `{"query": "mutation {subscribeBlog (blog_author_id:%d user_id:%d subscribe:%t) {id subscribers {id name last_name second_name email position personal_number photo login_ad}}}"}`
