package hugo

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

type Page struct {
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

func NewPage(title string) Page {
	return Page{title: title}
}

func (m Page) Content() (interface{}, error) {
	panic("implement me")
}

func (m Page) Plain() string {
	panic("implement me")
}

func (m Page) PlainWords() []string {
	panic("implement me")
}

func (m Page) Summary() template.HTML {
	panic("implement me")
}

func (m Page) Truncated() bool {
	panic("implement me")
}

func (m Page) FuzzyWordCount() int {
	panic("implement me")
}

func (m Page) WordCount() int {
	panic("implement me")
}

func (m Page) ReadingTime() int {
	panic("implement me")
}

func (m Page) Len() int {
	panic("implement me")
}

func (m Page) TableOfContents() template.HTML {
	panic("implement me")
}

func (m Page) RawContent() string {
	panic("implement me")
}

func (m Page) MediaType() media.Type {
	panic("implement me")
}

func (m Page) ResourceType() string {
	panic("implement me")
}

func (m Page) Permalink() string {
	panic("implement me")
}

func (m Page) RelPermalink() string {
	panic("implement me")
}

func (m Page) Name() string {
	panic("implement me")
}

func (m Page) Title() string {
	return m.title
}

func (m Page) Params() maps.Params {
	panic("implement me")
}

func (m Page) Data() interface{} {
	panic("implement me")
}

func (m Page) Date() time.Time {
	panic("implement me")
}

func (m Page) Lastmod() time.Time {
	panic("implement me")
}

func (m Page) PublishDate() time.Time {
	panic("implement me")
}

func (m Page) ExpiryDate() time.Time {
	panic("implement me")
}

func (m Page) Aliases() []string {
	panic("implement me")
}

func (m Page) BundleType() files.ContentClass {
	panic("implement me")
}

func (m Page) Description() string {
	panic("implement me")
}

func (m Page) Draft() bool {
	panic("implement me")
}

func (m Page) IsHome() bool {
	panic("implement me")
}

func (m Page) Keywords() []string {
	panic("implement me")
}

func (m Page) Kind() string {
	panic("implement me")
}

func (m Page) Layout() string {
	panic("implement me")
}

func (m Page) LinkTitle() string {
	panic("implement me")
}

func (m Page) IsNode() bool {
	panic("implement me")
}

func (m Page) IsPage() bool {
	panic("implement me")
}

func (m Page) Param(key interface{}) (interface{}, error) {
	panic("implement me")
}

func (m Page) Path() string {
	panic("implement me")
}

func (m Page) Slug() string {
	panic("implement me")
}

func (m Page) Lang() string {
	panic("implement me")
}

func (m Page) IsSection() bool {
	panic("implement me")
}

func (m Page) Section() string {
	panic("implement me")
}

func (m Page) SectionsEntries() []string {
	panic("implement me")
}

func (m Page) SectionsPath() string {
	panic("implement me")
}

func (m Page) Sitemap() config.Sitemap {
	panic("implement me")
}

func (m Page) Type() string {
	panic("implement me")
}

func (m Page) Weight() int {
	panic("implement me")
}

func (m Page) Language() *langs.Language {
	panic("implement me")
}

func (m Page) File() source.File {
	panic("implement me")
}

func (m Page) GitInfo() *gitmap.GitInfo {
	panic("implement me")
}

func (m Page) OutputFormats() page.OutputFormats {
	panic("implement me")
}

func (m Page) AlternativeOutputFormats() page.OutputFormats {
	panic("implement me")
}

func (m Page) Pages() page.Pages {
	panic("implement me")
}

func (m Page) RegularPages() page.Pages {
	panic("implement me")
}

func (m Page) RegularPagesRecursive() page.Pages {
	panic("implement me")
}

func (m Page) Resources() resource.Resources {
	panic("implement me")
}

func (m Page) IsAncestor(other interface{}) (bool, error) {
	panic("implement me")
}

func (m Page) CurrentSection() page.Page {
	panic("implement me")
}

func (m Page) IsDescendant(other interface{}) (bool, error) {
	panic("implement me")
}

func (m Page) FirstSection() page.Page {
	panic("implement me")
}

func (m Page) InSection(other interface{}) (bool, error) {
	panic("implement me")
}

func (m Page) Parent() page.Page {
	panic("implement me")
}

func (m Page) Sections() page.Pages {
	panic("implement me")
}

func (m Page) Page() page.Page {
	panic("implement me")
}

func (m Page) NextInSection() page.Page {
	panic("implement me")
}

func (m Page) PrevInSection() page.Page {
	panic("implement me")
}

func (m Page) Render(layout ...string) (template.HTML, error) {
	panic("implement me")
}

func (m Page) RenderString(args ...interface{}) (template.HTML, error) {
	panic("implement me")
}

func (m Page) Paginator(options ...interface{}) (*page.Pager, error) {
	panic("implement me")
}

func (m Page) Paginate(seq interface{}, options ...interface{}) (*page.Pager, error) {
	panic("implement me")
}

func (m Page) Next() page.Page {
	panic("implement me")
}

func (m Page) Prev() page.Page {
	panic("implement me")
}

func (m Page) PrevPage() page.Page {
	panic("implement me")
}

func (m Page) NextPage() page.Page {
	panic("implement me")
}

func (m Page) Menus() navigation.PageMenus {
	panic("implement me")
}

func (m Page) HasMenuCurrent(menuID string, me *navigation.MenuEntry) bool {
	panic("implement me")
}

func (m Page) IsMenuCurrent(menuID string, inme *navigation.MenuEntry) bool {
	panic("implement me")
}

func (m Page) Author() page.Author {
	panic("implement me")
}

func (m Page) Authors() page.AuthorList {
	panic("implement me")
}

func (m Page) GetPage(ref string) (page.Page, error) {
	panic("implement me")
}

func (m Page) Ref(argsm map[string]interface{}) (string, error) {
	panic("implement me")
}

func (m Page) RefFrom(argsm map[string]interface{}, source interface{}) (string, error) {
	panic("implement me")
}

func (m Page) RelRef(argsm map[string]interface{}) (string, error) {
	panic("implement me")
}

func (m Page) RelRefFrom(argsm map[string]interface{}, source interface{}) (string, error) {
	panic("implement me")
}

func (m Page) TranslationKey() string {
	panic("implement me")
}

func (m Page) IsTranslated() bool {
	panic("implement me")
}

func (m Page) AllTranslations() page.Pages {
	panic("implement me")
}

func (m Page) Translations() page.Pages {
	panic("implement me")
}

func (m Page) Site() page.Site {
	panic("implement me")
}

func (m Page) Sites() page.Sites {
	panic("implement me")
}

func (m Page) HasShortcode(name string) bool {
	panic("implement me")
}

func (m Page) Eq(other interface{}) bool {
	panic("implement me")
}

func (m Page) Scratch() *maps.Scratch {
	panic("implement me")
}

func (m Page) RelatedKeywords(cfg related.IndexConfig) ([]related.Keyword, error) {
	panic("implement me")
}

func (m Page) GetTerms(taxonomy string) page.Pages {
	panic("implement me")
}

func (m Page) Filename() string {
	panic("implement me")
}

func (m Page) Dir() string {
	panic("implement me")
}

func (m Page) Extension() string {
	panic("implement me")
}

func (m Page) Ext() string {
	panic("implement me")
}

func (m Page) LogicalName() string {
	panic("implement me")
}

func (m Page) BaseFileName() string {
	panic("implement me")
}

func (m Page) TranslationBaseName() string {
	panic("implement me")
}

func (m Page) ContentBaseName() string {
	panic("implement me")
}

func (m Page) UniqueID() string {
	panic("implement me")
}

func (m Page) FileInfo() hugofs.FileMetaInfo {
	panic("implement me")
}

func (m Page) IsDraft() bool {
	panic("implement me")
}

func (m Page) Hugo() hugo.Info {
	panic("implement me")
}

func (m Page) LanguagePrefix() string {
	panic("implement me")
}

func (m Page) GetParam(key string) interface{} {
	panic("implement me")
}

func (m Page) RSSLink() template.URL {
	panic("implement me")
}

func (m Page) URL() string {
	panic("implement me")
}
