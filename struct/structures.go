package structure

type WebText struct {
    Profile Profile
    InfoBio InfoBio
    NavItems []NavItems
    Experiences []Experience
    Projects []Project
    HomelabItems []Homelab
}

type Profile struct{
    Name string
    JobTitle string
    Bio string
    ImgURL string
}

type InfoBio struct {
    Title string
    Introduction string
    SkillCategories []SkillCategory
}

type SkillCategory struct {
    Title   string
    Skills  []Skill
}

type Skill struct {
    Name string
    IconURL string
}

type NavItems struct {
    Name string
    Id string
}

type Experience struct {
    Name string
    Date string
    Position string
    Description string
    PresentDesc string
}

type Project struct{
    Title string
    Description string
    Languages []string
    Link string
}

type Homelab struct {
    Title string
    Description string
    ImgURL string
}

