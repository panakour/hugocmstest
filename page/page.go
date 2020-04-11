package page

import (
	"github.com/bep/gitmap"
	"github.com/gohugoio/hugo/common/hugo"
	"github.com/gohugoio/hugo/common/maps"
	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugofs/files"
	"github.com/gohugoio/hugo/langs"
	"github.com/gohugoio/hugo/media"
	"github.com/gohugoio/hugo/navigation"
	"github.com/gohugoio/hugo/related"
	"github.com/gohugoio/hugo/resources/page"
	"github.com/gohugoio/hugo/resources/resource"
	"github.com/gohugoio/hugo/source"
	"html/template"
	"time"
)

type Myypage struct {
	description string
	title       string
	linkTitle   string

	section string

	content string

	fuzzyWordCount int

	path string

	slug string

	// Dates
	date       time.Time
	lastMod    time.Time
	expiryDate time.Time
	pubDate    time.Time

	weight int

	params map[string]interface{}
	data   map[string]interface{}

	//file source.File
}


func NewMyyPage(title string) Myypage {
	return Myypage{title: title}
}

func (m Myypage) Content() (interface{}, error) {
	panic("implement me")
}

func (m Myypage) Plain() string {
	panic("implement me")
}

func (m Myypage) PlainWords() []string {
	panic("implement me")
}

func (m Myypage) Summary() template.HTML {
	panic("implement me")
}

func (m Myypage) Truncated() bool {
	panic("implement me")
}

func (m Myypage) FuzzyWordCount() int {
	panic("implement me")
}

func (m Myypage) WordCount() int {
	panic("implement me")
}

func (m Myypage) ReadingTime() int {
	panic("implement me")
}

func (m Myypage) Len() int {
	panic("implement me")
}

func (m Myypage) TableOfContents() template.HTML {
	panic("implement me")
}

func (m Myypage) RawContent() string {
	panic("implement me")
}

func (m Myypage) MediaType() media.Type {
	panic("implement me")
}

func (m Myypage) ResourceType() string {
	panic("implement me")
}

func (m Myypage) Permalink() string {
	panic("implement me")
}

func (m Myypage) RelPermalink() string {
	panic("implement me")
}

func (m Myypage) Name() string {
	panic("implement me")
}

func (m Myypage) Title() string {
	return m.title
}

func (m Myypage) Params() maps.Params {
	panic("implement me")
}

func (m Myypage) Data() interface{} {
	panic("implement me")
}

func (m Myypage) Date() time.Time {
	panic("implement me")
}

func (m Myypage) Lastmod() time.Time {
	panic("implement me")
}

func (m Myypage) PublishDate() time.Time {
	panic("implement me")
}

func (m Myypage) ExpiryDate() time.Time {
	panic("implement me")
}

func (m Myypage) Aliases() []string {
	panic("implement me")
}

func (m Myypage) BundleType() files.ContentClass {
	panic("implement me")
}

func (m Myypage) Description() string {
	panic("implement me")
}

func (m Myypage) Draft() bool {
	panic("implement me")
}

func (m Myypage) IsHome() bool {
	panic("implement me")
}

func (m Myypage) Keywords() []string {
	panic("implement me")
}

func (m Myypage) Kind() string {
	panic("implement me")
}

func (m Myypage) Layout() string {
	panic("implement me")
}

func (m Myypage) LinkTitle() string {
	panic("implement me")
}

func (m Myypage) IsNode() bool {
	panic("implement me")
}

func (m Myypage) IsPage() bool {
	panic("implement me")
}

func (m Myypage) Param(key interface{}) (interface{}, error) {
	panic("implement me")
}

func (m Myypage) Path() string {
	panic("implement me")
}

func (m Myypage) Slug() string {
	panic("implement me")
}

func (m Myypage) Lang() string {
	panic("implement me")
}

func (m Myypage) IsSection() bool {
	panic("implement me")
}

func (m Myypage) Section() string {
	panic("implement me")
}

func (m Myypage) SectionsEntries() []string {
	panic("implement me")
}

func (m Myypage) SectionsPath() string {
	panic("implement me")
}

func (m Myypage) Sitemap() config.Sitemap {
	panic("implement me")
}

func (m Myypage) Type() string {
	panic("implement me")
}

func (m Myypage) Weight() int {
	panic("implement me")
}

func (m Myypage) Language() *langs.Language {
	panic("implement me")
}

func (m Myypage) File() source.File {
	panic("implement me")
}

func (m Myypage) GitInfo() *gitmap.GitInfo {
	panic("implement me")
}

func (m Myypage) OutputFormats() page.OutputFormats {
	panic("implement me")
}

func (m Myypage) AlternativeOutputFormats() page.OutputFormats {
	panic("implement me")
}

func (m Myypage) Pages() page.Pages {
	panic("implement me")
}

func (m Myypage) RegularPages() page.Pages {
	panic("implement me")
}

func (m Myypage) RegularPagesRecursive() page.Pages {
	panic("implement me")
}

func (m Myypage) Resources() resource.Resources {
	panic("implement me")
}

func (m Myypage) IsAncestor(other interface{}) (bool, error) {
	panic("implement me")
}

func (m Myypage) CurrentSection() page.Page {
	panic("implement me")
}

func (m Myypage) IsDescendant(other interface{}) (bool, error) {
	panic("implement me")
}

func (m Myypage) FirstSection() page.Page {
	panic("implement me")
}

func (m Myypage) InSection(other interface{}) (bool, error) {
	panic("implement me")
}

func (m Myypage) Parent() page.Page {
	panic("implement me")
}

func (m Myypage) Sections() page.Pages {
	panic("implement me")
}

func (m Myypage) Page() page.Page {
	panic("implement me")
}

func (m Myypage) NextInSection() page.Page {
	panic("implement me")
}

func (m Myypage) PrevInSection() page.Page {
	panic("implement me")
}

func (m Myypage) Render(layout ...string) (template.HTML, error) {
	panic("implement me")
}

func (m Myypage) RenderString(args ...interface{}) (template.HTML, error) {
	panic("implement me")
}

func (m Myypage) Paginator(options ...interface{}) (*page.Pager, error) {
	panic("implement me")
}

func (m Myypage) Paginate(seq interface{}, options ...interface{}) (*page.Pager, error) {
	panic("implement me")
}

func (m Myypage) Next() page.Page {
	panic("implement me")
}

func (m Myypage) Prev() page.Page {
	panic("implement me")
}

func (m Myypage) PrevPage() page.Page {
	panic("implement me")
}

func (m Myypage) NextPage() page.Page {
	panic("implement me")
}

func (m Myypage) Menus() navigation.PageMenus {
	panic("implement me")
}

func (m Myypage) HasMenuCurrent(menuID string, me *navigation.MenuEntry) bool {
	panic("implement me")
}

func (m Myypage) IsMenuCurrent(menuID string, inme *navigation.MenuEntry) bool {
	panic("implement me")
}

func (m Myypage) Author() page.Author {
	panic("implement me")
}

func (m Myypage) Authors() page.AuthorList {
	panic("implement me")
}

func (m Myypage) GetPage(ref string) (page.Page, error) {
	panic("implement me")
}

func (m Myypage) Ref(argsm map[string]interface{}) (string, error) {
	panic("implement me")
}

func (m Myypage) RefFrom(argsm map[string]interface{}, source interface{}) (string, error) {
	panic("implement me")
}

func (m Myypage) RelRef(argsm map[string]interface{}) (string, error) {
	panic("implement me")
}

func (m Myypage) RelRefFrom(argsm map[string]interface{}, source interface{}) (string, error) {
	panic("implement me")
}

func (m Myypage) TranslationKey() string {
	panic("implement me")
}

func (m Myypage) IsTranslated() bool {
	panic("implement me")
}

func (m Myypage) AllTranslations() page.Pages {
	panic("implement me")
}

func (m Myypage) Translations() page.Pages {
	panic("implement me")
}

func (m Myypage) Site() page.Site {
	panic("implement me")
}

func (m Myypage) Sites() page.Sites {
	panic("implement me")
}

func (m Myypage) HasShortcode(name string) bool {
	panic("implement me")
}

func (m Myypage) Eq(other interface{}) bool {
	panic("implement me")
}

func (m Myypage) Scratch() *maps.Scratch {
	panic("implement me")
}

func (m Myypage) RelatedKeywords(cfg related.IndexConfig) ([]related.Keyword, error) {
	panic("implement me")
}

func (m Myypage) GetTerms(taxonomy string) page.Pages {
	panic("implement me")
}

func (m Myypage) Filename() string {
	panic("implement me")
}

func (m Myypage) Dir() string {
	panic("implement me")
}

func (m Myypage) Extension() string {
	panic("implement me")
}

func (m Myypage) Ext() string {
	panic("implement me")
}

func (m Myypage) LogicalName() string {
	panic("implement me")
}

func (m Myypage) BaseFileName() string {
	panic("implement me")
}

func (m Myypage) TranslationBaseName() string {
	panic("implement me")
}

func (m Myypage) ContentBaseName() string {
	panic("implement me")
}

func (m Myypage) UniqueID() string {
	panic("implement me")
}

func (m Myypage) FileInfo() hugofs.FileMetaInfo {
	panic("implement me")
}

func (m Myypage) IsDraft() bool {
	panic("implement me")
}

func (m Myypage) Hugo() hugo.Info {
	panic("implement me")
}

func (m Myypage) LanguagePrefix() string {
	panic("implement me")
}

func (m Myypage) GetParam(key string) interface{} {
	panic("implement me")
}

func (m Myypage) RSSLink() template.URL {
	panic("implement me")
}

func (m Myypage) URL() string {
	panic("implement me")
}
