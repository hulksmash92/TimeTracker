import { ComponentFixture, TestBed } from '@angular/core/testing';
import { DebugElement } from '@angular/core';
import { RouterTestingModule } from '@angular/router/testing';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatListItem, MatListModule } from '@angular/material/list';

import { SettingsComponent } from './settings.component';
import { UserProfileFormComponent } from './components/user-profile-form/user-profile-form.component';
import { UserDeleteFormComponent } from './components/user-delete-form/user-delete-form.component';
import { UserDeleteConfirmComponent } from './components/user-delete-confirm/user-delete-confirm.component';
import { AuthService } from 'src/app/services/auth/auth.service';
import { UserService } from 'src/app/services/user/user.service';
import { MockAuthService, MockUserService, TestingPage } from 'src/app/testing';

/**
 * Represents the template of the SettingsComponent using the Page model pattern
 * Makes extracting elements from the template cleaner in our UI tests
 */
class SettingsComponentTestingPage extends TestingPage {
  /**
   * List items in the side nav
   */
  get listItems(): DebugElement[] {
    return this.queryAllByDirective(MatListItem);
  }

  /**
   * The form to delete a user
   */
  get deleteForm(): DebugElement { 
    return this.queryByDirective(UserDeleteFormComponent); 
  }

  /**
   * The form to update a user's profile/details
   */
  get profileForm(): DebugElement { 
    return this.queryByDirective(UserProfileFormComponent); 
  }

  constructor(fixture: ComponentFixture<SettingsComponent>) {
    super(fixture);
  }
}

describe('SettingsComponent', () => {
  let component: SettingsComponent;
  let fixture: ComponentFixture<SettingsComponent>;
  let page: SettingsComponentTestingPage;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        SettingsComponent,
        UserProfileFormComponent,
        UserDeleteFormComponent,
        UserDeleteConfirmComponent
      ],
      imports: [
        NoopAnimationsModule,
        ReactiveFormsModule,
        RouterTestingModule.withRoutes([]),
        MatButtonModule,
        MatCardModule,
        MatDialogModule,
        MatFormFieldModule,
        MatIconModule,
        MatInputModule,
        MatListModule,
      ],
      providers: [
        { provide: AuthService, useClass: MockAuthService },
        { provide: UserService, useClass: MockUserService }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SettingsComponent);
    page = new SettingsComponentTestingPage(fixture);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  describe('#changeVisibleSection()', () => {
    it('should change #visibleSection to account', () => {
      // set visibleSection to nothing
      component.visibleSection = null;

      // call the func with `account` passed in
      component.changeVisibleSection('account');

      // assert that visibleSection is `account`
      expect(component.visibleSection).toEqual('account');
    });

    it('should change #visibleSection to profile', () => {
      // set visibleSection to nothing
      component.visibleSection = null;

      // call the func with `profile` passed in
      component.changeVisibleSection('profile');

      // assert that visibleSection is `profile`
      expect(component.visibleSection).toEqual('profile');
    });
  });

  describe('template', () => {
    const navActiveClass = 'settings-nav-active';

    it('should show the profile from when #visibleSection=profile', () => {
      // Set the visibleSection property to profile
      component.visibleSection = 'profile';

      // Manually run change detection on the component fixture so the change is picked up
      fixture.detectChanges();

      // Assert that the profile form is shown and not the delete form
      expect(page.profileForm).toBeTruthy('Profile form not displayed');
      expect(page.deleteForm).toBeFalsy('Delete form is displayed, should be showing profile form only');

      // Assert that the the profile button is showing as active
      expect(page.listItems[0].nativeElement.classList).toContain(navActiveClass, 'Profile button should be active');
      expect(page.listItems[1].nativeElement.classList).not.toContain(navActiveClass, 'Account button should not be active');
    });

    it('should show the account from when #visibleSection=account', () => {
      // Set the visibleSection property to account
      component.visibleSection = 'account';

      // Manually run change detection on the component fixture so the change is picked up
      fixture.detectChanges();

      // Assert that the delete form is shown and not the profile form
      expect(page.deleteForm).toBeTruthy('Delete form not displayed');
      expect(page.profileForm).toBeFalsy('Profile form is displayed, should be showing profile form only');

      // Assert that the the account button is showing as active
      expect(page.listItems[0].nativeElement.classList).not.toContain(navActiveClass, 'Profile button should not be active');
      expect(page.listItems[1].nativeElement.classList).toContain(navActiveClass, 'Account button should be active');
    });

    it('profile button click event handler should trigger #changeVisibleSection() with `profile` as param value', () => {
      // spy on the changeVisibleSection() function and replace its implementation with a mock
      const changeVisibleSectionSpy = spyOn(component, 'changeVisibleSection').and.callFake((section: string) => {});

      // trigger the click event handler on the profile button (the first list item in the template)
      page.listItems[0].triggerEventHandler('click', null);

      // assert that the changeVisibleSection() function was called with profile 
      // by checking the spy we created earlier
      expect(changeVisibleSectionSpy).toHaveBeenCalledWith('profile');
    });

  });

});
