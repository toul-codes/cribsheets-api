import { Component, OnInit } from '@angular/core';
import { Account } from '../../../models/account';
import { CribsheetsApiService } from '../../services/cribsheets-api.service';


@Component({
  selector: 'sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  private account : Account;
  public showMsg: boolean;

  constructor(private cribsheetsApiService: CribsheetsApiService){
    this.showMsg = false;
  }
  ngOnInit() {
  }

  
  saveNewUser(username, password) {
    this.account = new Account(username, password)
    this.cribsheetsApiService.signUp(this.account).subscribe(res => {
      this.showMsg = true;
    }, err => {
      this.showMsg = false;
    })
  }

 

}
