package moviebuff

// Certification contains a readable code as well as UUID along with whether the certification indicates that the movie is safe for children and the country it is applicable for.
// A movie has multiple certifications, one for each country it is released in.
type Certification struct {
	//Indicates whether this certification represents safe for children
	ChildSafe bool `json:"childSafe"`
	// UUID of this certification
	UUID string `json:"uuid"`
	// Readable code for this certification
	Code string `json:"code"`
	// Country of this certification
	Country Country `json:"country"`
}

//Country has country information as available on Qube Wire Cinemas
type Country struct {
	//Country name as on Qube Wire Cinemas
	Name string `json:"name"`
	//ISO 2-digit Country Code
	Code string `json:"code"`
	//Country UUID as on Qube Wire Cinemas
	UUID string `json:"uuid"`
}
