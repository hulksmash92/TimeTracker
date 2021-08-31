import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { AuthComponent } from './auth.component';
import { AuthService } from 'src/app/services/auth/auth.service';
import { MockAuthService } from 'src/app/testing';

describe('AuthComponent', () => {
  let component: AuthComponent;
  let fixture: ComponentFixture<AuthComponent>;
  let authService: AuthService;

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
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  describe('#ngOnDestroy()', () => {
    let unsubscribeSpy: jasmine.Spy;

    beforeEach(() => {
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

  // TODO: test getAccessToken()

});
