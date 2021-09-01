import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { AuthService } from './auth.service';
import { WindowService } from 'src/app/services/window/window.service';
import { MockWindowService } from 'src/app/testing';

describe('AuthService', () => {
  let service: AuthService;
  let windowService: WindowService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        HttpClientTestingModule,
        RouterTestingModule.withRoutes([])
      ],
      providers: [
        { provide: WindowService, useClass: MockWindowService },
      ]
    });
    service = TestBed.inject(AuthService);
    windowService = TestBed.inject(WindowService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
