import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NewCribsheetComponent } from './new-cribsheet.component';

describe('NewCribsheetComponent', () => {
  let component: NewCribsheetComponent;
  let fixture: ComponentFixture<NewCribsheetComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NewCribsheetComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NewCribsheetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
