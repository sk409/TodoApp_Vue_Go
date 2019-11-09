import Model from "./model.js";

export default class Todo extends Model {

    static fetch(params, handler) {
        Model.fetch(Todo.basePath(), params, handler);
    }

    static basePath() {
        return "todos";
    }

    constructor(id = null, title = "", content = "", limit = "") {
        super(id)
        this.title = title;
        this.content = content;
        this.limit = limit;
    }

    basePath() {
        return Todo.basePath();
    }

    params() {
        return {
            title: this.title,
            content: this.content,
            limit: this.limit
        }
    }

    get title() {
        return this._title;
    }

    set title(value) {
        this._title = value;
    }

    get content() {
        return this._content;
    }

    set content(value) {
        this._content = value;
    }

    get limit() {
        return this._limit;
    }

    set limit(value) {
        this._limit = value;
    }

}