import { Component, OnInit } from '@angular/core';

import { User } from 'src/app/models/user';
import { UserService } from 'src/app/services/user/user.service';

@Component({
  selector: 'user-menu',
  templateUrl: './user-menu.component.html',
  styleUrls: ['./user-menu.component.scss']
})
export class UserMenuComponent implements OnInit {
  user: User;

  constructor(private readonly userService: UserService) { }

  ngOnInit(): void {
    this.userService.get().subscribe((res: any) => {
      this.user = res;
    });
  }

}
