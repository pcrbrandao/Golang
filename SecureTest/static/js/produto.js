/**
 * Created by pcrbrandao on 11/02/17.
 */
$('[name=tamanho]').on('input', function(){
    $('[name=valortamanho]').val(this.value);
});