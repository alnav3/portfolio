package models

type Profile struct{
    Name string
    JobTitle string
    Bio string
}

type ProjectBio struct{
    Title string
    Description string
}

type Project struct{
    Title string
    Description string
    Languages []string
}

type InfoBio struct {
    Introduction string
    Skills       []SkillCategory
}

type SkillCategory struct {
    Title   string
    Skills  []Skill
}

type Skill struct {
    Name string
    IconURL string
}


