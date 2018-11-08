import { TestBed, inject } from '@angular/core/testing';

import { CribsheetsApiService } from './cribsheets-api.service';

describe('CribsheetsApiService', () => {
  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [CribsheetsApiService]
    });
  });

  it('should be created', inject([CribsheetsApiService], (service: CribsheetsApiService) => {
    expect(service).toBeTruthy();
  }));
});
