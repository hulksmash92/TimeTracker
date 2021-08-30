import { Location as NgLocationService } from '@angular/common';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { of } from 'rxjs';

import { MaterialModule } from 'src/app/modules/material/material.module';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';
import { MockAuthService, MockUserService } from 'src/app/testing';
import { AvatarModule } from 'src/app/components/avatar';
import { UserMenuComponent } from './user-menu.component';

describe('UserMenuComponent', () => {
  let component: UserMenuComponent;
  let fixture: ComponentFixture<UserMenuComponent>;
  let authService: AuthService;
  let locationService: NgLocationService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserMenuComponent ],
      imports: [
        RouterTestingModule,
        MaterialModule,
        AvatarModule
      ],
      providers: [
        { provide: UserService, useClass: MockUserService },
        { provide: AuthService, useClass: MockAuthService }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserMenuComponent);
    authService = TestBed.inject(AuthService);
    locationService = TestBed.inject(NgLocationService);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  describe('#handleGitHubLogin()', () => {
    let gitHubUrlSpy: jasmine.Spy;
    let goSpy: jasmine.Spy;

    beforeEach(() => {
      // Spy on our AuthService.gitHubUrl() method
      gitHubUrlSpy = spyOn(authService, 'gitHubUrl');

      // spy on the NgLocationService.go() method 
      // and replace the implementation with a mock func
      const mockGoFn = (path: string, query?: string, state?: any) => {};
      goSpy = spyOn(locationService, 'go').and.callFake(mockGoFn);
    });

    it('should call locationService.go() when AuthService.getHubUrl() returns a truthy value', () => {
      // Valid mock url
      const mockUrl = 'https://github.com/login/oauth/authorize?clientId=gh123456&scopes=user:email';

      // mock the return value of the AuthService.gitHubUrl() 
      // method using our spy by chaining `and.returnValue()`
      // wrapping it in rxjs of() to turn the value into an Observable
      gitHubUrlSpy.and.returnValue(of(mockUrl));

      // Call the #handleGitHubLogin() method we're testing
      component.handleGitHubLogin();

      // Assert that AuthService.gitHubUrl() was called
      expect(gitHubUrlSpy).toHaveBeenCalled();

      // Assert that NgLocationService.go() was called with the mockUrl value
      expect(goSpy).toHaveBeenCalledWith(mockUrl);
    });

    it('should not call locationService.go() when AuthService.getHubUrl() returns a falsy value', () => {
      // mock a null return value of the AuthService.gitHubUrl() 
      gitHubUrlSpy.and.returnValue(of(null));

      // Call the #handleGitHubLogin() method we're testing
      component.handleGitHubLogin();

      // Assert that AuthService.gitHubUrl() was called
      expect(gitHubUrlSpy).toHaveBeenCalled();

      // Assert that NgLocationService.go() was not called
      expect(goSpy).not.toHaveBeenCalledWith();
    });
  });
  
});
