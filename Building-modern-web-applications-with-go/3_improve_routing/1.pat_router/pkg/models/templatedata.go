package models

//this struct holds any possible kind of data that we could send to the template
// holds data set from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{} //we're not sure what kind of data that may be, so using interface
	CSRFToken string //cross site request forgery token = random string that takes care for authorization (we will use nusof package)
	Flash string // messaged
	Warning string //warnings
	Error string
}

