import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ListCribsheetsComponent } from './list-cribsheets.component';

describe('ListCribsheetsComponent', () => {
  let component: ListCribsheetsComponent;
  let fixture: ComponentFixture<ListCribsheetsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ListCribsheetsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ListCribsheetsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
