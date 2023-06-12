package controllers

import (
	"encoding/json"
	_"fmt"
	"net/http"
   _ "main/models"
)


func AllProducts(w http.ResponseWriter, r *http.Request) {
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
