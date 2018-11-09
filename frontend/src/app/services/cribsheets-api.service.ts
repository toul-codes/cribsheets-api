import { Injectable, Inject } from '@angular/core';
import { Http, Headers } from '@angular/http';
import 'rxjs/add/operator/map';
import { environment } from '../../environments/environment';
import { Cribsheet } from '../../models/cribsheet';
import {LOCAL_STORAGE, WebStorageService} from 'angular-webstorage-service';

@Injectable()
export class CribsheetsApiService {

  constructor(private http:Http, @Inject(LOCAL_STORAGE) private storage: WebStorageService) { }

  findAll(){
    return this.http
    .get(environment.api, {headers: this.getHeaders()})
    .map(res => {
      return res.json()
    })
  }

  insert(cribsheet: Cribsheet){
    return this.http
      .post(environment.api, JSON.stringify(cribsheet), {headers: this.getHeaders()})
      .map(res => {
        return res
      })
}
getHeaders(){
  let headers = new Headers()
  headers.append('Authorization', this.storage.get("COGNITO_TOKEN"))
  return headers
}
}
