import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';

import { of } from 'rxjs';

import { AuthComponent } from './auth.component';
import { AuthService } from 'src/app/services/auth/auth.service';
import { MockAuthService } from 'src/app/testing';
import { User } from 'src/app/models/user';

// Create a mock user object for testing
const mockUser: User = {
  id: 1,
  githubUserId: 'hulksmash92',
  name: 'Alex Deakins',
  email: 'user@example.com',
  created: new Date(),
  updated: new Date(),
  avatar: null,
  organisations: [],
  apiClients: []
};

describe('AuthComponent', () => {
  let component: AuthComponent;
  let fixture: ComponentFixture<AuthComponent>;
  let authService: AuthService;
  let router: Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AuthComponent ],
      imports: [
        RouterTestingModule.withRoutes([]),
      ],
      providers: [
        { provide: AuthService, useClass: MockAuthService }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AuthComponent);
    authService = TestBed.inject(AuthService);
    router = TestBed.inject(Router);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  describe('#ngOnDestroy()', () => {
    let unsubscribeSpy: jasmine.Spy;

    beforeEach(() => {
      // Spy on the unsubscribe method on the route query param subscription
      // so we can check if its been called correctly
      unsubscribeSpy = spyOn(component.routeQuerySub, 'unsubscribe');
    });

    it('should call #routeQuerySub.unsubscribe() when #routeQuerySub.closed is false', () => {
      // Set the closed prop to false on the #routeQuerySub subscription
      component.routeQuerySub.closed = false;

      // Call ngOnDestroy()
      component.ngOnDestroy();

      // assert that the unsubscribe method was called
      expect(unsubscribeSpy).toHaveBeenCalled();
    });

    it('should not call #routeQuerySub.unsubscribe() when #routeQuerySub.closed is true', () => {
      // Set the closed prop to true on the #routeQuerySub subscription
      // to indicate that the subscription is now closed
      component.routeQuerySub.closed = true;

      // Call ngOnDestroy()
      component.ngOnDestroy();

      // assert that the unsubscribe method was called
      expect(unsubscribeSpy).not.toHaveBeenCalled();
    });
  });

  describe('#getAccessToken()', () => {
    let loginGitHubSpy: jasmine.Spy;
    let navigateSpy: jasmine.Spy;

    beforeEach(() => {
      // Spy on the loginGitHub method on AuthService so we can check if its been 
      // called correctly and mock the returned values for certain edge cases
      loginGitHubSpy = spyOn(authService, 'loginGitHub');

      // Spy on the Router.navigate() method and replace its
      // implementation with a mock func
      navigateSpy = spyOn(router, 'navigate').and.callFake((commands: any[], extras?: any) => Promise.resolve(true));
    });

    it('should call #authService.loginGitHub() with session code and set authService.user to response', () => {
      const mockSessionCode = 'abc1234def56';

      // initialise the user property in authService to null
      authService.user = null;
      
      // mock the return value of the AuthService.loginGitHub() to a fake user value
      loginGitHubSpy.and.returnValue(of(mockUser));

      // call the getAccessToken() method for testing
      component.getAccessToken(mockSessionCode);

      // assert that the service method called with the correct value
      expect(loginGitHubSpy).toHaveBeenCalledWith(mockSessionCode);

      // assert that user has been set correctly
      expect(authService.user).toEqual(mockUser);

      // assert that the user is being redirected to /time
      expect(navigateSpy).toHaveBeenCalledWith(['/time']);
    });

    it('should not call #router.navigate() when #authService.loginGitHub() response is falsy', () => {
      // mock the return value of the AuthService.loginGitHub() to a null value
      loginGitHubSpy.and.returnValue(of(null));

      // call the getAccessToken() method for testing
      component.getAccessToken('abc1234def56');

      // assert router navigation has not been called
      expect(navigateSpy).not.toHaveBeenCalled();
    });
  });

});
