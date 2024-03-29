import { Component, OnInit } from '@angular/core';
import { Cribsheet } from '../../../models/cribsheet';
import { CribsheetsApiService } from '../../services/cribsheets-api.service';

@Component({
  selector: 'new-cribsheet',
  templateUrl: './new-cribsheet.component.html',
  styleUrls: ['./new-cribsheet.component.css']
})
export class NewCribsheetComponent implements OnInit {

  private cribsheet : Cribsheet;
  public showMsg: boolean;

  constructor(private cribsheetsApiService: CribsheetsApiService) {
    this.showMsg = false;
  }

  ngOnInit() {
  }

  save(professor, coursenumber, description, filename) {
    this.cribsheet = new Cribsheet(professor, coursenumber, description, filename)
    this.cribsheetsApiService.insert(this.cribsheet).subscribe(res => {
      this.showMsg = true;
    }, err => {
      this.showMsg = false;
    })
}

}

