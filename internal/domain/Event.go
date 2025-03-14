package domain

type GitHubEvent struct {
    Repo     string 
    Type     string 
    Action   string 
    PRURL    string 
    CommitID string 
    User     string 
    Branch   string
}