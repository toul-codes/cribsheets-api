import { Component, OnInit, Input } from '@angular/core';
import { Cribsheet } from '../../../models/cribsheet';
@Component({
  selector: 'cribsheet-item',
  templateUrl: './cribsheet-item.component.html',
  styleUrls: ['./cribsheet-item.component.css']
})
export class CribsheetItemComponent implements OnInit {
  @Input()
  public cribsheet: Cribsheet;
  constructor() { }

  ngOnInit() {
  }

}
