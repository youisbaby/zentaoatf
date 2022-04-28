import { Mutation, Action } from 'vuex';
import { StoreModuleType } from "@/utils/store";
import { ResponseData } from '@/utils/request';

import {
    list, get, extract, create, update, remove, loadChildren, updateCode, syncFromZentao
} from './service';
import {ScriptFileNotExist} from "@/utils/const";

export interface ScriptData {
    list: [];
    detail: any;

    currWorkspace: any
    queryParams: any;
}

export interface ModuleType extends StoreModuleType<ScriptData> {
    state: ScriptData;
    mutations: {
        setList: Mutation<ScriptData>;
        setItem: Mutation<ScriptData>;
        setWorkspace: Mutation<ScriptData>;
        setQueryParams: Mutation<ScriptData>;
    };
    actions: {
        listScript: Action<ScriptData, ScriptData>;
        getScript: Action<ScriptData, ScriptData>;
        loadChildren: Action<ScriptData, ScriptData>;
        syncFromZentao: Action<ScriptData, ScriptData>;
        extractScript: Action<ScriptData, ScriptData>;
        changeWorkspace: Action<ScriptData, ScriptData>;

        createScript: Action<ScriptData, ScriptData>;
        updateScript: Action<ScriptData, ScriptData>;
        deleteScript: Action<ScriptData, ScriptData>;
        updateCode: Action<ScriptData, ScriptData>;
    };
}
const initState: ScriptData = {
    list: [],
    detail: null,

    currWorkspace: {id: 0, type: 'ztf'},
    queryParams: {},
};

const StoreModel: ModuleType = {
    namespaced: true,
    name: 'Script',
    state: {
        ...initState
    },
    mutations: {
        setList(state, payload) {
            state.list = payload;
        },
        setItem(state, payload) {
            state.detail = payload;
        },
        setWorkspace(state, payload) {
            state.currWorkspace = payload;
        },
        setQueryParams(state, payload) {
            state.queryParams = payload;
        },
    },
    actions: {
        async listScript({ commit }, playload: any ) {
            const response: ResponseData = await list(playload);
            const { data } = response;
            commit('setList', [data]);

            commit('setQueryParams', playload);
            return true;
        },
        async loadChildren({ commit }, treeNode: any ) {
            console.log('load node children', treeNode.dataRef.workspaceType)
            if (treeNode.dataRef.workspaceType === 'ztf')
                return true

            loadChildren(treeNode.dataRef.path, treeNode.dataRef.workspaceId).then((json) => {
                treeNode.dataRef.children = json.data
                return true;
            })
        },

        async getScript({ commit }, script: any ) {
            if (!script || script.type !== 'file') {
                commit('setItem', null);
                return true;
            }

            if (script.path.indexOf('zentao-') === 0) {
                commit('setItem', {id: script.caseId, workspaceId: script.workspaceId, code: ScriptFileNotExist});
                return true;
            }

            const response: ResponseData = await get(script.path, script.workspaceId);
            commit('setItem', response.data);
            return true;
        },

        async syncFromZentao({ commit, dispatch, state }, payload: any ) {
            const resp = await syncFromZentao(payload)
            if (resp.code === 0) {
                await dispatch('listScript', state.queryParams)

                if (resp.code === 0 && resp.data.length === 1) {
                    const getResp = await get(resp.data[0], payload.workspaceId);
                    commit('setItem', getResp.data);
                } else {
                    commit('setItem', null);
                }
            }

            return resp
        },

        async extractScript({ commit }, script: any ) {
            if (!script.path) return true

            const response: ResponseData = await extract(script.path, script.workspaceId)
            const { data } = response
            commit('setItem', data.script)

            return data.done
        },

        async createScript({ commit }, payload: any) {
            try {
                await create(payload);
                return true;
            } catch (error) {
                return false;
            }
        },
        async updateScript({ commit }, payload: any ) {
            try {
                const { id, ...params } = payload;
                await update(id, { ...params });
                return true;
            } catch (error) {
                return false;
            }
        },

        async deleteScript({ commit }, payload: number ) {
            try {
                await remove(payload);
                return true;
            } catch (error) {
                return false;
            }
        },

        async updateCode({ commit }, payload: any ) {
            try {
                await updateCode(payload);
                return true;
            } catch (error) {
                return false;
            }
        },

        async changeWorkspace({ commit }, payload: any ) {
            commit('setWorkspace', payload);
            return true;
        },
    }
};

export default StoreModel;
