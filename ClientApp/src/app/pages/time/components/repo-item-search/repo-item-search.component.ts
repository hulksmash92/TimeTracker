import { Component, EventEmitter, Output, ViewChild, Input } from '@angular/core';
import { FormControl } from '@angular/forms';
import { MatListOption, MatSelectionList, MatSelectionListChange } from '@angular/material/list';

import { RepoItem } from 'src/app/models/repo-item';
import { RepoSearchResult } from 'src/app/models/repos';
import { RepoService } from 'src/app/services/repo/repo.service';

export interface RepoItemType {
  name: string;
  label: string;
}

@Component({
  selector: 'repo-item-search',
  templateUrl: './repo-item-search.component.html',
  styleUrls: ['./repo-item-search.component.scss']
})
export class RepoItemSearchComponent {
  @Input() repo: RepoSearchResult;
  @Input() source: string = 'GitHub';
  @Output() repoItemsSelected: EventEmitter<RepoItem[]> = new EventEmitter<RepoItem[]>();
  @ViewChild(MatSelectionList, { static: true }) matSelectionList: MatSelectionList;
  readonly repoItemTypes: RepoItemType[] = [
    { name: 'branch', label: 'Branches' },
    { name: 'commit', label: 'Commits' },
  ];
  repoItems: any[] = [];
  repoItemType: FormControl = new FormControl();
  itemFrom: FormControl = new FormControl();
  itemTo: FormControl = new FormControl();

  get itemType(): string {
    return this.repoItemType.value;
  }

  constructor(private readonly repoService: RepoService) { }

  getRepoItems(): void {
    if (!!this.repo && !!this.itemType) {
      const from: Date = this.itemFrom.value;
      const to: Date = this.itemFrom.value;

      this.repoService.getGitHubRepoItems(this.repo.owner, this.repo.name, this.itemType, from, to)
        .subscribe((res: any[]) => {
          this.repoItems = res;
        });
    }
  }

  emitItems(selectChange: MatSelectionListChange): void {
    if (selectChange.options.length > 0) {
      const items = selectChange.options.map(o => this.parseSelectOption(o));
      this.repoItemsSelected.emit(items);
      this.matSelectionList.deselectAll();
    }
  }

  parseSelectOption(option: MatListOption): RepoItem {
    const item = option.value;
    const itemType = this.itemType;
    let description: string;
    let itemIdSource: string;

    switch (itemType) {
      case 'branch':
        itemIdSource = item.commit.node_id;
        description = `${item.name}, commit: ${item.commit.commit.message}`;
        break;
      case 'commit':
        itemIdSource = item.node_id;
        description = item.message;
        break;
    }

    const newRepoItem: RepoItem = {
      itemIdSource,
      itemType,
      description,
      source: this.source,
      repoName: this.repo.fullName
    };
    return newRepoItem;
  }

}
