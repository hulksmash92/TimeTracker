import { Component, EventEmitter, OnInit, Output, OnDestroy, ViewChild  } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { MatMenuTrigger } from '@angular/material/menu';

import { RepoSearchResult } from 'src/app/models/repos';
import { RepoService } from 'src/app/services/repo/repo.service';

@Component({
  selector: 'repo-search',
  templateUrl: './repo-search.component.html',
  styleUrls: ['./repo-search.component.scss']
})
export class RepoSearchComponent implements OnInit, OnDestroy {
  @Output() repoSelected: EventEmitter<any> = new EventEmitter<any>();
  @ViewChild(MatMenuTrigger, {static: true}) matMenuTrigger: MatMenuTrigger;
  searchFc: FormControl = new FormControl(null, [Validators.minLength(3)]);
  result: RepoSearchResult[];

  get searchValue(): string {
    return this.searchFc.value;
  }

  constructor(private readonly repoService: RepoService) { }

  ngOnInit(): void {
    this.matMenuTrigger.menuOpened.subscribe({
      next: () => {
        if (!this.result) {
          this.matMenuTrigger.closeMenu();
        }
      }
    });
  }

  ngOnDestroy(): void {
    this.matMenuTrigger.closeMenu();
  }

  search(): void {
    if (this.searchFc.valid) {
      this.repoService.searchGitHub(this.searchValue.trim())
        .subscribe((res: RepoSearchResult[]) => {
          this.result = res || [];
          this.matMenuTrigger.openMenu();
        });
    }
  }

}
