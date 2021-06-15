import { RepoItem } from './repo-item';
import { Tag } from './tag';

export interface TimeEntry {
    id?: number;
    userId?: number;
    organisationId?: number;
    comments: string;
    created?: string;
    updated?: string;
    value: number;
    valueType: string;
    tags?: Tag[];
    repoItems?: RepoItem[];
}