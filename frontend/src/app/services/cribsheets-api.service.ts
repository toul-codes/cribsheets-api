import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';
import { Cribsheet } from '../../models/cribsheet';

@Injectable()
export class CribsheetsApiService {

  constructor(private http:Http) { }

  findAll(){
    return this.http
    .get(environment.api)
    .map(res => {
      return res.json()
    })
  }

  insert(cribsheet: Cribsheet){
    return this.http
      .post(environment.api, JSON.stringify(cribsheet))
      .map(res => {
        return res
      })
}
}
