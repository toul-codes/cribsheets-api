import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpModule } from '@angular/http';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { CribsheetsApiService } from './services/cribsheets-api.service';
import { AppComponent } from './app.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { CribsheetItemComponent } from './components/cribsheet-item/cribsheet-item.component';
import { ListCribsheetsComponent } from './components/list-cribsheets/list-cribsheets.component';
import { NewCribsheetComponent } from './components/new-cribsheet/new-cribsheet.component';
import { CognitoService } from './services/cognito.service';

import { StorageServiceModule} from 'angular-webstorage-service';
import { SignUpComponent } from './components/sign-up/sign-up.component';


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    CribsheetItemComponent,
    ListCribsheetsComponent,
    NewCribsheetComponent,
    SignUpComponent,

  ],
  imports: [
    NgbModule.forRoot(),
    BrowserModule,
    HttpModule,
    StorageServiceModule

  ],
  providers: [
    CribsheetsApiService,
    CognitoService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
