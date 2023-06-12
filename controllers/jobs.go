package controllers;

import (
	"encoding/json"
	_"fmt"
	"net/http"
   _ "main/models"
)

type Response struct {
	Message string `json:"message"`
	Response string `json:"response"`
}

type Bhangudatas struct {
	Id int `json:"id"`
	TrackId *int64 `json:"track_id"`
	CustomerName *string `json:"customer_name,omitempty"`
	Date *string `json:"date,omitempty"`
	Datetimestamp *string `json:"datetimestamp"`
	Myob *string `json:"myob,omitempty"`
	Installer *string `json:"installer,omitempty"`
	AInstallerId *int64 `json:"a_installer_id,omitempty"`
	Address *string `json:"address,omitempty"`
	PoFile *string `json:"po_file"`
	Po *string `json:"po"`
	Subtotal *string `json:"subtotal"`
	Qty *float64 `json:"qty,omitempty"`
	WallWrap *float64 `json:"wall_wrap"`
	WallBats *float64 `json:"wall_bats"`
	WallBatsR27 *int64 `json:"wall_bats_r2_7"`
	Celling *float64 `json:"celling"`
	StringCelling *float64 `json:"string_celling"`
	R50 *int64 `json:"r5_0"`
	Dishing *int64 `json:"dishing"`
	Cutting *float64 `json:"cutting"`
	Fireseal *int64 `json:"fireseal"`
	Fixup *int64 `json:"fixup"`
	WindowTape *int64 `json:"window_tape"`
	Tape *int64 `json:"tape"`
	DampCourse *int64 `json:"damp_course"`
	Infill *int64 `json:"infill"`
	EmPrice *float64 `json:"em_price"`
	Productcode *string `json:"productcode"`
	Description *string `json:"description,omitempty"`
	Notes *string `json:"notes,omitempty"`
	ScopeOfWork *string `json:"scope_of_work,omitempty"`
	UnitPrice *string `json:"unit_price"`
	UnitPriceGst *string `json:"unit_price_gst"`
	Delivery *string `json:"delivery,omitempty"`
	FletcherPo *string `json:"fletcher_po,omitempty"`
	Comments *string `json:"comments,omitempty"`
	Pid *int64 `json:"pid"`
	Status *string `json:"status"`
	InstallerNote *string `json:"installer_note"`
	AdminNote *string `json:"admin_note"`
	AdminComment *string `json:"admin_comment"`
	InvoiceReady *int64 `json:"invoice_ready"`
	JsaLink *string `json:"jsa_link"`
	AInstaller *string `json:"a_installer,omitempty"`
	JobTimestamp *int64 `json:"job_timestamp"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	NeedInfill string `json:"need_infill"`
	NeedLintel string `json:"need_lintel"`
}



func Allobs(w http.ResponseWriter, r *http.Request) {
	var bhangudatas []Bhangudatas
	//var allMaps []map[string]*int64erface{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	result, err := db.Query("SELECT * FROM bhangudatas  ORDER BY id DESC");

	if err != nil {
		panic(err.Error())
	  }
	  defer result.Close()
	  for result.Next() {
		var bhangudata Bhangudatas
		err := result.Scan(&bhangudata.Id,&bhangudata.TrackId,&bhangudata.CustomerName,&bhangudata.Date,&bhangudata.Datetimestamp,&bhangudata.Myob,&bhangudata.Installer,&bhangudata.AInstallerId,&bhangudata.Address,&bhangudata.PoFile,&bhangudata.Po,&bhangudata.Subtotal,&bhangudata.Qty,&bhangudata.WallWrap,&bhangudata.WallBats,&bhangudata.WallBatsR27,&bhangudata.Celling,&bhangudata.StringCelling,&bhangudata.R50,&bhangudata.Dishing,&bhangudata.Cutting,&bhangudata.Fireseal,&bhangudata.Fixup,&bhangudata.WindowTape,&bhangudata.Tape,&bhangudata.DampCourse,&bhangudata.Infill,&bhangudata.EmPrice,&bhangudata.Productcode,&bhangudata.Description,&bhangudata.Notes,&bhangudata.ScopeOfWork,&bhangudata.UnitPrice,&bhangudata.UnitPriceGst,&bhangudata.Delivery,&bhangudata.FletcherPo,&bhangudata.Comments,&bhangudata.Pid,&bhangudata.Status,&bhangudata.InstallerNote,&bhangudata.AdminNote,&bhangudata.AdminComment,&bhangudata.InvoiceReady,&bhangudata.JsaLink,&bhangudata.AInstaller,&bhangudata.JobTimestamp,&bhangudata.CreatedAt,&bhangudata.UpdatedAt,&bhangudata.NeedInfill,&bhangudata.NeedLintel)
		if err != nil {
		  panic(err.Error())
		}
		bhangudatas = append(bhangudatas, bhangudata)
	  }
	  json.NewEncoder(w).Encode(bhangudatas)
}
