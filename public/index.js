const app = new Vue({
    el: '#app',
    data: {
        categories: [],
        currentCategory: '',
        files: []
    },
    created() {
        this.getAllCategory()
            .then(res => {
                this.categories = res;

                return res;
            })
            .then(res => {
                if (res.length > 0) {
                    this.currentCategory = res[0];
                    this.getFileByCategory(this.currentCategory)
                        .then(res => {
                            this.files = res;
                        });
                }
            })
    },
    mounted() {

    },
    methods: {
        getAllCategory() {
            return fetch('/categories')
                .then(function (response) {
                    return response.json();
                });
        },
        getFileByCategory(category) {
            return fetch('/file?category=' + category)
                .then(function (response) {
                    return response.json();
                });
        },
        clickCategory(category) {
            this.currentCategory = category;
            this.getFileByCategory(category)
                .then(res => {
                    this.files = res;
                });
        },
        clickFile(file) {

        }
    }
})