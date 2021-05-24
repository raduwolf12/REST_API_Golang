import { TestBed } from '@angular/core/testing';

import { RestRequestService } from './rest-request.service';

describe('RestRequestService', () => {
  let service: RestRequestService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RestRequestService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
