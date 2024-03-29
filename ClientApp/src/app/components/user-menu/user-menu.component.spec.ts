import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';

import { of } from 'rxjs';

import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';
import { MockAuthService, MockUserService } from 'src/app/testing';
import { AvatarModule } from 'src/app/components/avatar';
import { UserMenuComponent } from './user-menu.component';
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

describe('UserMenuComponent', () => {
  let component: UserMenuComponent;
  let fixture: ComponentFixture<UserMenuComponent>;
  let authService: AuthService;
  let userService: UserService;
  let router: Router;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserMenuComponent ],
      imports: [
        RouterTestingModule.withRoutes([]),
        MatButtonModule,
        MatIconModule,
        MatMenuModule,
        AvatarModule
      ],
      providers: [
        // Mock our UserService and AuthService dependencies to avoid having 
        // to import the HttpClient module to simplify our tests
        { provide: UserService, useClass: MockUserService },
        { provide: AuthService, useClass: MockAuthService }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserMenuComponent);
    authService = TestBed.inject(AuthService);
    userService = TestBed.inject(UserService);
    router = TestBed.inject(Router);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('#user getter should return #authService.user value', () => {
    // test null user value
    authService.user = null;
    expect(component.user).toBeNull();

    // test undefined user value
    authService.user = undefined;
    expect(component.user).toBeUndefined();

    // test a valid user object using the mockUser value above
    authService.user = mockUser;
    expect(component.user).toEqual(mockUser);
  });

  describe('#ngOnInit()', () => {
    let userGetSpy: jasmine.Spy;
    let urlSpy: jasmine.Spy;

    beforeEach(() => {
      // Spy on the UserService.get() method
      userGetSpy = spyOn(userService, 'get');

      // Spy on the Router.url getter property
      urlSpy = spyOnProperty(router, 'url', 'get');
    });

    it('should set AuthService.user to response from UserService.get() when Router.url is not `auth`', () => {
      // Instantiate our AuthService.user value to null
      authService.user = null;

      // Mock Router.url to be `time` using our urlSpy
      urlSpy.and.returnValue('time');

      // Mock the UserService.get() method return value with the mockUser above with our userGetSpy
      userGetSpy.and.returnValue(of(mockUser));

      // Call ngOnInit() so we can test the implementation
      component.ngOnInit();

      // Assert that UserService.get() was called
      expect(userGetSpy).toHaveBeenCalled();

      // Assert that AuthService.user was set correctly
      expect(authService.user).toEqual(mockUser);
    });

    it('should not callUserService.get() when Router.url is `auth`', () => {
      // Mock Router.url to be `time` using our urlSpy
      urlSpy.and.returnValue('auth');

      // Call ngOnInit() so we can test the implementation
      component.ngOnInit();

      // Assert that UserService.get() was not called
      expect(userGetSpy).not.toHaveBeenCalled();
    });
  });

  // TODO: Add tests for component template
  
});
