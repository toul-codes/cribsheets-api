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


@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    CribsheetItemComponent,
    ListCribsheetsComponent,
    NewCribsheetComponent
  ],
  imports: [
    NgbModule.forRoot(),
    BrowserModule,
    HttpModule,

  ],
  providers: [
    CribsheetsApiService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
