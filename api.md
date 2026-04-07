# Posts

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostNewResponse">PostNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostGetResponse">PostGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostUpdateResponse">PostUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostListResponse">PostListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostBulkNewResponse">PostBulkNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostRetryResponse">PostRetryResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostUnpublishResponse">PostUnpublishResponse</a>

Methods:

- <code title="post /v1/posts">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostNewParams">PostNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostNewResponse">PostNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/posts/{id}">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostGetResponse">PostGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/posts/{id}">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostUpdateParams">PostUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostUpdateResponse">PostUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/posts">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostListParams">PostListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostListResponse">PostListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/posts/{id}">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/posts/bulk">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.BulkNew">BulkNew</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostBulkNewParams">PostBulkNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostBulkNewResponse">PostBulkNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/posts/{id}/retry">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.Retry">Retry</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostRetryResponse">PostRetryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/posts/{id}/unpublish">client.Posts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostService.Unpublish">Unpublish</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostUnpublishParams">PostUnpublishParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostUnpublishResponse">PostUnpublishResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logs

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogGetResponse">PostLogGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogListResponse">PostLogListResponse</a>

Methods:

- <code title="get /v1/posts/{id}/logs">client.Posts.Logs.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogGetResponse">PostLogGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/posts/logs">client.Posts.Logs.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogListParams">PostLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#PostLogListResponse">PostLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGetResponse">AccountGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountUpdateResponse">AccountUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountListResponse">AccountListResponse</a>

Methods:

- <code title="get /v1/accounts/{id}">client.Accounts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/accounts/{id}">client.Accounts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountUpdateParams">AccountUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountUpdateResponse">AccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts">client.Accounts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/accounts/{id}">client.Accounts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Health

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthGetResponse">AccountHealthGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthListResponse">AccountHealthListResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/health">client.Accounts.Health.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthGetResponse">AccountHealthGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/accounts/health">client.Accounts.Health.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthListParams">AccountHealthListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountHealthListResponse">AccountHealthListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RedditFlairs

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditFlairGetResponse">AccountRedditFlairGetResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/reddit-flairs">client.Accounts.RedditFlairs.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditFlairService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditFlairGetParams">AccountRedditFlairGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditFlairGetResponse">AccountRedditFlairGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FacebookPages

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageGetResponse">AccountFacebookPageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageSetDefaultResponse">AccountFacebookPageSetDefaultResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/facebook-pages">client.Accounts.FacebookPages.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageGetResponse">AccountFacebookPageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/{id}/facebook-pages">client.Accounts.FacebookPages.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageService.SetDefault">SetDefault</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageSetDefaultParams">AccountFacebookPageSetDefaultParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountFacebookPageSetDefaultResponse">AccountFacebookPageSetDefaultResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## LinkedinOrganizations

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationGetResponse">AccountLinkedinOrganizationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationSwitchTypeResponse">AccountLinkedinOrganizationSwitchTypeResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/linkedin-organizations">client.Accounts.LinkedinOrganizations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationGetResponse">AccountLinkedinOrganizationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/{id}/linkedin-organizations">client.Accounts.LinkedinOrganizations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationService.SwitchType">SwitchType</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationSwitchTypeParams">AccountLinkedinOrganizationSwitchTypeParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountLinkedinOrganizationSwitchTypeResponse">AccountLinkedinOrganizationSwitchTypeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PinterestBoards

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardGetResponse">AccountPinterestBoardGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardSetDefaultResponse">AccountPinterestBoardSetDefaultResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/pinterest-boards">client.Accounts.PinterestBoards.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardGetResponse">AccountPinterestBoardGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/{id}/pinterest-boards">client.Accounts.PinterestBoards.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardService.SetDefault">SetDefault</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardSetDefaultParams">AccountPinterestBoardSetDefaultParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountPinterestBoardSetDefaultResponse">AccountPinterestBoardSetDefaultResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RedditSubreddits

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditGetResponse">AccountRedditSubredditGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditSetDefaultResponse">AccountRedditSubredditSetDefaultResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/reddit-subreddits">client.Accounts.RedditSubreddits.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditGetResponse">AccountRedditSubredditGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/{id}/reddit-subreddits">client.Accounts.RedditSubreddits.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditService.SetDefault">SetDefault</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditSetDefaultParams">AccountRedditSubredditSetDefaultParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountRedditSubredditSetDefaultResponse">AccountRedditSubredditSetDefaultResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## GmbLocations

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationGetResponse">AccountGmbLocationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationSetDefaultResponse">AccountGmbLocationSetDefaultResponse</a>

Methods:

- <code title="get /v1/accounts/{id}/gmb-locations">client.Accounts.GmbLocations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationGetResponse">AccountGmbLocationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/accounts/{id}/gmb-locations">client.Accounts.GmbLocations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationService.SetDefault">SetDefault</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationSetDefaultParams">AccountGmbLocationSetDefaultParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AccountGmbLocationSetDefaultResponse">AccountGmbLocationSetDefaultResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Media

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaGetResponse">MediaGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaGetPresignURLResponse">MediaGetPresignURLResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaUploadResponse">MediaUploadResponse</a>

Methods:

- <code title="get /v1/media/{id}">client.Media.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaGetResponse">MediaGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/media/{id}">client.Media.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/media/presign">client.Media.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaService.GetPresignURL">GetPresignURL</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaGetPresignURLParams">MediaGetPresignURLParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaGetPresignURLResponse">MediaGetPresignURLResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/media/upload">client.Media.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaUploadParams">MediaUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#MediaUploadResponse">MediaUploadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookUpdateResponse">WebhookUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookListResponse">WebhookListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookListLogsResponse">WebhookListLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookSendTestResponse">WebhookSendTestResponse</a>

Methods:

- <code title="post /v1/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v1/webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookUpdateParams">WebhookUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookUpdateResponse">WebhookUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookListParams">WebhookListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookListResponse">WebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/webhooks/{id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1/webhooks/logs">client.Webhooks.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookService.ListLogs">ListLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookListLogsParams">WebhookListLogsParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookListLogsResponse">WebhookListLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/webhooks/test">client.Webhooks.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookService.SendTest">SendTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookSendTestParams">WebhookSendTestParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WebhookSendTestResponse">WebhookSendTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyNewResponse">APIKeyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyListResponse">APIKeyListResponse</a>

Methods:

- <code title="post /v1/api-keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyNewParams">APIKeyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyNewResponse">APIKeyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/api-keys">client.APIKeys.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyListParams">APIKeyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyListResponse">APIKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/api-keys/{id}">client.APIKeys.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#APIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Usage

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#UsageGetResponse">UsageGetResponse</a>

Methods:

- <code title="get /v1/usage">client.Usage.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#UsageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#UsageGetResponse">UsageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AccountGroups

# Connect

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectCompleteOAuthCallbackResponse">ConnectCompleteOAuthCallbackResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectNewBlueskyConnectionResponse">ConnectNewBlueskyConnectionResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFetchPendingDataResponse">ConnectFetchPendingDataResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectStartOAuthFlowResponse">ConnectStartOAuthFlowResponse</a>

Methods:

- <code title="post /v1/connect/{platform}">client.Connect.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectService.CompleteOAuthCallback">CompleteOAuthCallback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, platform <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectCompleteOAuthCallbackParamsPlatform">ConnectCompleteOAuthCallbackParamsPlatform</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectCompleteOAuthCallbackParams">ConnectCompleteOAuthCallbackParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectCompleteOAuthCallbackResponse">ConnectCompleteOAuthCallbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/bluesky">client.Connect.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectService.NewBlueskyConnection">NewBlueskyConnection</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectNewBlueskyConnectionParams">ConnectNewBlueskyConnectionParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectNewBlueskyConnectionResponse">ConnectNewBlueskyConnectionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/connect/pending-data">client.Connect.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectService.FetchPendingData">FetchPendingData</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFetchPendingDataParams">ConnectFetchPendingDataParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFetchPendingDataResponse">ConnectFetchPendingDataResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/connect/{platform}">client.Connect.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectService.StartOAuthFlow">StartOAuthFlow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, platform <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectStartOAuthFlowParamsPlatform">ConnectStartOAuthFlowParamsPlatform</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectStartOAuthFlowParams">ConnectStartOAuthFlowParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectStartOAuthFlowResponse">ConnectStartOAuthFlowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Telegram

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramConnectDirectlyResponse">ConnectTelegramConnectDirectlyResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramInitiateConnectionResponse">ConnectTelegramInitiateConnectionResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramPollConnectionStatusResponse">ConnectTelegramPollConnectionStatusResponse</a>

Methods:

- <code title="post /v1/connect/telegram/direct">client.Connect.Telegram.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramService.ConnectDirectly">ConnectDirectly</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramConnectDirectlyParams">ConnectTelegramConnectDirectlyParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramConnectDirectlyResponse">ConnectTelegramConnectDirectlyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/telegram">client.Connect.Telegram.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramService.InitiateConnection">InitiateConnection</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramInitiateConnectionResponse">ConnectTelegramInitiateConnectionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/connect/telegram">client.Connect.Telegram.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramService.PollConnectionStatus">PollConnectionStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramPollConnectionStatusParams">ConnectTelegramPollConnectionStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectTelegramPollConnectionStatusResponse">ConnectTelegramPollConnectionStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Whatsapp

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappCompleteEmbeddedSignupResponse">ConnectWhatsappCompleteEmbeddedSignupResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappConnectViaCredentialsResponse">ConnectWhatsappConnectViaCredentialsResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappGetSDKConfigResponse">ConnectWhatsappGetSDKConfigResponse</a>

Methods:

- <code title="post /v1/connect/whatsapp/embedded-signup">client.Connect.Whatsapp.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappService.CompleteEmbeddedSignup">CompleteEmbeddedSignup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappCompleteEmbeddedSignupParams">ConnectWhatsappCompleteEmbeddedSignupParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappCompleteEmbeddedSignupResponse">ConnectWhatsappCompleteEmbeddedSignupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/whatsapp/credentials">client.Connect.Whatsapp.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappService.ConnectViaCredentials">ConnectViaCredentials</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappConnectViaCredentialsParams">ConnectWhatsappConnectViaCredentialsParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappConnectViaCredentialsResponse">ConnectWhatsappConnectViaCredentialsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/connect/whatsapp/sdk-config">client.Connect.Whatsapp.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappService.GetSDKConfig">GetSDKConfig</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectWhatsappGetSDKConfigResponse">ConnectWhatsappGetSDKConfigResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Facebook

### Pages

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageListResponse">ConnectFacebookPageListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageSelectResponse">ConnectFacebookPageSelectResponse</a>

Methods:

- <code title="get /v1/connect/facebook/pages">client.Connect.Facebook.Pages.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageListResponse">ConnectFacebookPageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/facebook/pages">client.Connect.Facebook.Pages.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageService.Select">Select</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageSelectParams">ConnectFacebookPageSelectParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectFacebookPageSelectResponse">ConnectFacebookPageSelectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Linkedin

### Organizations

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationListResponse">ConnectLinkedinOrganizationListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationSelectResponse">ConnectLinkedinOrganizationSelectResponse</a>

Methods:

- <code title="get /v1/connect/linkedin/organizations">client.Connect.Linkedin.Organizations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationListResponse">ConnectLinkedinOrganizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/linkedin/organizations">client.Connect.Linkedin.Organizations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationService.Select">Select</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationSelectParams">ConnectLinkedinOrganizationSelectParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectLinkedinOrganizationSelectResponse">ConnectLinkedinOrganizationSelectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Pinterest

### Boards

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardListResponse">ConnectPinterestBoardListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardSelectResponse">ConnectPinterestBoardSelectResponse</a>

Methods:

- <code title="get /v1/connect/pinterest/boards">client.Connect.Pinterest.Boards.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardListResponse">ConnectPinterestBoardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/pinterest/boards">client.Connect.Pinterest.Boards.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardService.Select">Select</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardSelectParams">ConnectPinterestBoardSelectParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectPinterestBoardSelectResponse">ConnectPinterestBoardSelectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Googlebusiness

### Locations

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationListResponse">ConnectGooglebusinessLocationListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationSelectResponse">ConnectGooglebusinessLocationSelectResponse</a>

Methods:

- <code title="get /v1/connect/googlebusiness/locations">client.Connect.Googlebusiness.Locations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationListResponse">ConnectGooglebusinessLocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/googlebusiness/locations">client.Connect.Googlebusiness.Locations.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationService.Select">Select</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationSelectParams">ConnectGooglebusinessLocationSelectParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectGooglebusinessLocationSelectResponse">ConnectGooglebusinessLocationSelectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Snapchat

### Profiles

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileListResponse">ConnectSnapchatProfileListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileSelectResponse">ConnectSnapchatProfileSelectResponse</a>

Methods:

- <code title="get /v1/connect/snapchat/profiles">client.Connect.Snapchat.Profiles.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileListResponse">ConnectSnapchatProfileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/connect/snapchat/profiles">client.Connect.Snapchat.Profiles.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileService.Select">Select</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileSelectParams">ConnectSnapchatProfileSelectParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectSnapchatProfileSelectResponse">ConnectSnapchatProfileSelectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Connections

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectionListLogsResponse">ConnectionListLogsResponse</a>

Methods:

- <code title="get /v1/connections/logs">client.Connections.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectionService.ListLogs">ListLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectionListLogsParams">ConnectionListLogsParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ConnectionListLogsResponse">ConnectionListLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Analytics

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetResponse">AnalyticsGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetBestTimeResponse">AnalyticsGetBestTimeResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetContentDecayResponse">AnalyticsGetContentDecayResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetPostTimelineResponse">AnalyticsGetPostTimelineResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetPostingFrequencyResponse">AnalyticsGetPostingFrequencyResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsListDailyMetricsResponse">AnalyticsListDailyMetricsResponse</a>

Methods:

- <code title="get /v1/analytics">client.Analytics.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetParams">AnalyticsGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetResponse">AnalyticsGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/analytics/best-time">client.Analytics.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsService.GetBestTime">GetBestTime</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetBestTimeParams">AnalyticsGetBestTimeParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetBestTimeResponse">AnalyticsGetBestTimeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/analytics/content-decay">client.Analytics.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsService.GetContentDecay">GetContentDecay</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetContentDecayParams">AnalyticsGetContentDecayParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetContentDecayResponse">AnalyticsGetContentDecayResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/analytics/post-timeline">client.Analytics.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsService.GetPostTimeline">GetPostTimeline</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetPostTimelineParams">AnalyticsGetPostTimelineParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetPostTimelineResponse">AnalyticsGetPostTimelineResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/analytics/posting-frequency">client.Analytics.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsService.GetPostingFrequency">GetPostingFrequency</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetPostingFrequencyParams">AnalyticsGetPostingFrequencyParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsGetPostingFrequencyResponse">AnalyticsGetPostingFrequencyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/analytics/daily-metrics">client.Analytics.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsService.ListDailyMetrics">ListDailyMetrics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsListDailyMetricsParams">AnalyticsListDailyMetricsParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsListDailyMetricsResponse">AnalyticsListDailyMetricsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Youtube

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsYoutubeGetDailyViewsResponse">AnalyticsYoutubeGetDailyViewsResponse</a>

Methods:

- <code title="get /v1/analytics/youtube/daily-views">client.Analytics.Youtube.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsYoutubeService.GetDailyViews">GetDailyViews</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsYoutubeGetDailyViewsParams">AnalyticsYoutubeGetDailyViewsParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#AnalyticsYoutubeGetDailyViewsResponse">AnalyticsYoutubeGetDailyViewsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

## Validate

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateCheckPostLengthResponse">ToolValidateCheckPostLengthResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateGetSubredditResponse">ToolValidateGetSubredditResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateValidateMediaResponse">ToolValidateValidateMediaResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateValidatePostResponse">ToolValidateValidatePostResponse</a>

Methods:

- <code title="post /v1/tools/validate/post-length">client.Tools.Validate.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateService.CheckPostLength">CheckPostLength</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateCheckPostLengthParams">ToolValidateCheckPostLengthParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateCheckPostLengthResponse">ToolValidateCheckPostLengthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools/validate/subreddit">client.Tools.Validate.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateService.GetSubreddit">GetSubreddit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateGetSubredditParams">ToolValidateGetSubredditParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateGetSubredditResponse">ToolValidateGetSubredditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/validate/media">client.Tools.Validate.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateService.ValidateMedia">ValidateMedia</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateValidateMediaParams">ToolValidateValidateMediaParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateValidateMediaResponse">ToolValidateValidateMediaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/tools/validate/post">client.Tools.Validate.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateService.ValidatePost">ValidatePost</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateValidatePostParams">ToolValidateValidatePostParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolValidateValidatePostResponse">ToolValidateValidatePostResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Instagram

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolInstagramCheckHashtagSafetyResponse">ToolInstagramCheckHashtagSafetyResponse</a>

Methods:

- <code title="post /v1/tools/instagram/hashtag-checker">client.Tools.Instagram.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolInstagramService.CheckHashtagSafety">CheckHashtagSafety</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolInstagramCheckHashtagSafetyParams">ToolInstagramCheckHashtagSafetyParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#ToolInstagramCheckHashtagSafetyResponse">ToolInstagramCheckHashtagSafetyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Queue

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueGetNextSlotResponse">QueueGetNextSlotResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueuePreviewResponse">QueuePreviewResponse</a>

Methods:

- <code title="get /v1/queue/next-slot">client.Queue.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueService.GetNextSlot">GetNextSlot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueGetNextSlotResponse">QueueGetNextSlotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/queue/preview">client.Queue.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueService.Preview">Preview</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueuePreviewParams">QueuePreviewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueuePreviewResponse">QueuePreviewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Slots

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotNewResponse">QueueSlotNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotUpdateResponse">QueueSlotUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotListResponse">QueueSlotListResponse</a>

Methods:

- <code title="post /v1/queue/slots">client.Queue.Slots.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotNewParams">QueueSlotNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotNewResponse">QueueSlotNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/queue/slots">client.Queue.Slots.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotUpdateParams">QueueSlotUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotUpdateResponse">QueueSlotUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/queue/slots">client.Queue.Slots.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotListResponse">QueueSlotListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/queue/slots">client.Queue.Slots.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#QueueSlotService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Twitter

## Retweet

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetNewResponse">TwitterRetweetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetUndoResponse">TwitterRetweetUndoResponse</a>

Methods:

- <code title="post /v1/twitter/retweet">client.Twitter.Retweet.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetNewParams">TwitterRetweetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetNewResponse">TwitterRetweetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/twitter/retweet">client.Twitter.Retweet.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetService.Undo">Undo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetUndoParams">TwitterRetweetUndoParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterRetweetUndoResponse">TwitterRetweetUndoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Bookmark

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkNewResponse">TwitterBookmarkNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkRemoveResponse">TwitterBookmarkRemoveResponse</a>

Methods:

- <code title="post /v1/twitter/bookmark">client.Twitter.Bookmark.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkNewParams">TwitterBookmarkNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkNewResponse">TwitterBookmarkNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/twitter/bookmark">client.Twitter.Bookmark.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkRemoveParams">TwitterBookmarkRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterBookmarkRemoveResponse">TwitterBookmarkRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Follow

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowNewResponse">TwitterFollowNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowUnfollowResponse">TwitterFollowUnfollowResponse</a>

Methods:

- <code title="post /v1/twitter/follow">client.Twitter.Follow.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowNewParams">TwitterFollowNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowNewResponse">TwitterFollowNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/twitter/follow">client.Twitter.Follow.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowService.Unfollow">Unfollow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowUnfollowParams">TwitterFollowUnfollowParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#TwitterFollowUnfollowResponse">TwitterFollowUnfollowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inbox

## Comments

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentGetResponse">InboxCommentGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentListResponse">InboxCommentListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentDeleteResponse">InboxCommentDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentPrivateReplyResponse">InboxCommentPrivateReplyResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentReplyResponse">InboxCommentReplyResponse</a>

Methods:

- <code title="get /v1/inbox/comments/{post_id}">client.Inbox.Comments.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, postID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentGetParams">InboxCommentGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentGetResponse">InboxCommentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/inbox/comments">client.Inbox.Comments.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentListParams">InboxCommentListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentListResponse">InboxCommentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/inbox/comments/{comment_id}">client.Inbox.Comments.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, commentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentDeleteResponse">InboxCommentDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inbox/comments/{comment_id}/private-reply">client.Inbox.Comments.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentService.PrivateReply">PrivateReply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, commentID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentPrivateReplyParams">InboxCommentPrivateReplyParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentPrivateReplyResponse">InboxCommentPrivateReplyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/inbox/comments/{post_id}/reply">client.Inbox.Comments.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentService.Reply">Reply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, postID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentReplyParams">InboxCommentReplyParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentReplyResponse">InboxCommentReplyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Hide

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentHideNewResponse">InboxCommentHideNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentHideDeleteResponse">InboxCommentHideDeleteResponse</a>

Methods:

- <code title="post /v1/inbox/comments/{comment_id}/hide">client.Inbox.Comments.Hide.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentHideService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, commentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentHideNewResponse">InboxCommentHideNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/inbox/comments/{comment_id}/hide">client.Inbox.Comments.Hide.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentHideService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, commentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentHideDeleteResponse">InboxCommentHideDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Like

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentLikeNewResponse">InboxCommentLikeNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentLikeDeleteResponse">InboxCommentLikeDeleteResponse</a>

Methods:

- <code title="post /v1/inbox/comments/{comment_id}/like">client.Inbox.Comments.Like.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentLikeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, commentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentLikeNewResponse">InboxCommentLikeNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/inbox/comments/{comment_id}/like">client.Inbox.Comments.Like.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentLikeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, commentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxCommentLikeDeleteResponse">InboxCommentLikeDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

## Reviews

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewListResponse">InboxReviewListResponse</a>

Methods:

- <code title="get /v1/inbox/reviews">client.Inbox.Reviews.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewListParams">InboxReviewListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewListResponse">InboxReviewListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Reply

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyNewResponse">InboxReviewReplyNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyDeleteResponse">InboxReviewReplyDeleteResponse</a>

Methods:

- <code title="post /v1/inbox/reviews/{review_id}/reply">client.Inbox.Reviews.Reply.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reviewID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyNewParams">InboxReviewReplyNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyNewResponse">InboxReviewReplyNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/inbox/reviews/{review_id}/reply">client.Inbox.Reviews.Reply.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, reviewID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#InboxReviewReplyDeleteResponse">InboxReviewReplyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Reddit

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditGetFeedResponse">RedditGetFeedResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditSearchResponse">RedditSearchResponse</a>

Methods:

- <code title="get /v1/reddit/feed">client.Reddit.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditService.GetFeed">GetFeed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditGetFeedParams">RedditGetFeedParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditGetFeedResponse">RedditGetFeedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/reddit/search">client.Reddit.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditSearchParams">RedditSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#RedditSearchResponse">RedditSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Whatsapp

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBulkSendResponse">WhatsappBulkSendResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappListPhoneNumbersResponse">WhatsappListPhoneNumbersResponse</a>

Methods:

- <code title="post /v1/whatsapp/bulk-send">client.Whatsapp.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappService.BulkSend">BulkSend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBulkSendParams">WhatsappBulkSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBulkSendResponse">WhatsappBulkSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/whatsapp/phone-numbers">client.Whatsapp.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappService.ListPhoneNumbers">ListPhoneNumbers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappListPhoneNumbersParams">WhatsappListPhoneNumbersParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappListPhoneNumbersResponse">WhatsappListPhoneNumbersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Broadcasts

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastNewResponse">WhatsappBroadcastNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastGetResponse">WhatsappBroadcastGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastListResponse">WhatsappBroadcastListResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastScheduleResponse">WhatsappBroadcastScheduleResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastSendResponse">WhatsappBroadcastSendResponse</a>

Methods:

- <code title="post /v1/whatsapp/broadcasts">client.Whatsapp.Broadcasts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastNewParams">WhatsappBroadcastNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastNewResponse">WhatsappBroadcastNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/whatsapp/broadcasts/{broadcast_id}">client.Whatsapp.Broadcasts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastGetResponse">WhatsappBroadcastGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/whatsapp/broadcasts">client.Whatsapp.Broadcasts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastListParams">WhatsappBroadcastListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastListResponse">WhatsappBroadcastListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/whatsapp/broadcasts/{broadcast_id}">client.Whatsapp.Broadcasts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/whatsapp/broadcasts/{broadcast_id}/schedule">client.Whatsapp.Broadcasts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastService.Schedule">Schedule</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastScheduleResponse">WhatsappBroadcastScheduleResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/whatsapp/broadcasts/{broadcast_id}/send">client.Whatsapp.Broadcasts.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, broadcastID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBroadcastSendResponse">WhatsappBroadcastSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Templates

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateNewResponse">WhatsappTemplateNewResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateGetResponse">WhatsappTemplateGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateListResponse">WhatsappTemplateListResponse</a>

Methods:

- <code title="post /v1/whatsapp/templates">client.Whatsapp.Templates.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateNewParams">WhatsappTemplateNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateNewResponse">WhatsappTemplateNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/whatsapp/templates/{template_name}">client.Whatsapp.Templates.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateGetParams">WhatsappTemplateGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateGetResponse">WhatsappTemplateGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/whatsapp/templates">client.Whatsapp.Templates.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateListParams">WhatsappTemplateListParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateListResponse">WhatsappTemplateListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/whatsapp/templates/{template_name}">client.Whatsapp.Templates.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappTemplateDeleteParams">WhatsappTemplateDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Contacts

## Groups

## BusinessProfile

Response Types:

- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileGetResponse">WhatsappBusinessProfileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileUpdateResponse">WhatsappBusinessProfileUpdateResponse</a>

Methods:

- <code title="get /v1/whatsapp/business-profile">client.Whatsapp.BusinessProfile.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileGetParams">WhatsappBusinessProfileGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileGetResponse">WhatsappBusinessProfileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/whatsapp/business-profile">client.Whatsapp.BusinessProfile.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileUpdateParams">WhatsappBusinessProfileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go">relaygo</a>.<a href="https://pkg.go.dev/github.com/relayapi-dev/relay-go#WhatsappBusinessProfileUpdateResponse">WhatsappBusinessProfileUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
