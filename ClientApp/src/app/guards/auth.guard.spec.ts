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
        // Replace the AuthService instance that is injected into the AuthGuard constructor
        // with an instance of our MockAuthService
        { provide: AuthService, useClass: MockAuthService }
      ]
    });

    // Get the instance of AuthGuard from our TestBed using `inject()`
    guard = TestBed.inject(AuthGuard);

    // Get the instance of AuthService created by the TestBed.
    // This will actually be an instance of MockAuthService as we've
    // replaced AuthService with that type in our providers array.
    authService = TestBed.inject(AuthService);
  });

  /**
   * Creates a new jasmine.Spy instance to spy on the AuthService.isAuthenticated() method 
   * with the return value set to the selected `returnValue` param value
   * 
   * @param returnValue boolean value to return
   * 
   * @returns the instance of jasmine.Spy created
   */
  function isAuthenticatedSpyFactory(returnValue: boolean): jasmine.Spy {
    return spyOn(authService, 'isAuthenticated').and.returnValue(Promise.resolve(returnValue));
  }

  it('#canActivate() should return value received from AuthService.isAuthenticated()', async () => {
    // Spy on our AuthService.isAuthenticated() method and mock it to return true
    const isAuthenticatedSpy = isAuthenticatedSpyFactory(true);

    // Grab the value inside the Promise returned by AuthGuard.canActivate() 
    const returnVal = await guard.canActivate({} as any, {} as any);

    // Assert that true was returned
    expect(returnVal).toBeTrue();

    // Assert that the AuthService.isAuthenticated() method was called
    expect(isAuthenticatedSpy).toHaveBeenCalled();
  });

  it('#canLoad() should return value received from AuthService.isAuthenticated()', async () => {
    // Spy on our AuthService.isAuthenticated() method and mock it to return false
    const isAuthenticatedSpy = isAuthenticatedSpyFactory(false);

    // Grab the value inside the Promise returned by AuthGuard.canLoad()
    const returnVal = await guard.canLoad({} as any, []);

    // Assert that false was returned
    expect(returnVal).toBeFalse();
    
    // Assert that the AuthService.isAuthenticated() method was called
    expect(isAuthenticatedSpy).toHaveBeenCalled();
  });
});
