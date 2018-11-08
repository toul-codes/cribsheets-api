import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CribsheetItemComponent } from './cribsheet-item.component';

describe('CribsheetItemComponent', () => {
  let component: CribsheetItemComponent;
  let fixture: ComponentFixture<CribsheetItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CribsheetItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CribsheetItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
