import { Component, OnInit } from '@angular/core';
import { Cribsheet } from '../../../models/cribsheet';
import { CribsheetsApiService } from '../../services/cribsheets-api.service';
@Component({
  selector: 'list-cribsheets',
  templateUrl: './list-cribsheets.component.html',
  styleUrls: ['./list-cribsheets.component.css']
})
export class ListCribsheetsComponent implements OnInit {
	
  public cribsheets: Cribsheet[];
  constructor(private cribsheetsApiService: CribsheetsApiService) {
    this.cribsheets = []
    this.cribsheetsApiService.findAll().subscribe(res => {
      res.forEach(cribsheet => {
        this.cribsheets.push(new Cribsheet(cribsheet.lastname, cribsheet.subject, cribsheet.type, cribsheet.university, cribsheet.year))
      })
    })
  } 
	
  ngOnInit() {
  }

}
