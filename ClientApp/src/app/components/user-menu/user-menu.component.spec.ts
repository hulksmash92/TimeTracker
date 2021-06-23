import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MaterialModule } from 'src/app/modules/material/material.module';
import { UserService } from 'src/app/services/user/user.service';
import { MockUserService } from 'src/app/testing/mock-user-service';
import { AvatarModule } from '../avatar/avatar.module';
import { UserMenuComponent } from './user-menu.component';

describe('UserMenuComponent', () => {
  let component: UserMenuComponent;
  let fixture: ComponentFixture<UserMenuComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UserMenuComponent ],
      imports: [
        MaterialModule,
        AvatarModule
      ],
      providers: [
        {
          provide: UserService,
          useClass: MockUserService
        }
      ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(UserMenuComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
