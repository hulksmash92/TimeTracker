import { TestBed } from '@angular/core/testing';

import { AuthGuard } from './auth.guard';
import { AuthService } from 'src/app/services/auth/auth.service';
import { MockAuthService } from 'src/app/testing';

describe('AuthGuard', () => {
  let guard: AuthGuard;
  let authService: AuthService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      providers: [
        { provide: AuthService, useClass: MockAuthService }
      ]
    });
    guard = TestBed.inject(AuthGuard);
    authService = TestBed.inject(AuthService);
  });

  function isAuthenticatedSpyFactory(returnVal: boolean): jasmine.Spy {
    return spyOn(authService, 'isAuthenticated').and.returnValue(Promise.resolve(returnVal));
  }

  it('#canActivate() should return value received from AuthService.isAuthenticated()', async () => {
    const isAuthenticatedSpy = isAuthenticatedSpyFactory(true);
    const returnVal = await guard.canActivate({} as any, {} as any);
    expect(returnVal).toBeTrue();
    expect(isAuthenticatedSpy).toHaveBeenCalled();
  });

  it('#canLoad() should return value received from AuthService.isAuthenticated()', async () => {
    const isAuthenticatedSpy = isAuthenticatedSpyFactory(false);
    const returnVal = await guard.canLoad({} as any, []);
    expect(returnVal).toBeFalse();
    expect(isAuthenticatedSpy).toHaveBeenCalled();
  });
});
