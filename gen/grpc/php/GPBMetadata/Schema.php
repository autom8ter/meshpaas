<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: schema.proto

namespace GPBMetadata;

class Schema
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\Struct::initOnce();
        \GPBMetadata\Google\Protobuf\Timestamp::initOnce();
        \GPBMetadata\Google\Protobuf\Any::initOnce();
        \GPBMetadata\Google\Protobuf\GPBEmpty::initOnce();
        \GPBMetadata\GithubCom\Mwitkow\GoProtoValidators\Validator::initOnce();
        $pool->internalAddGeneratedFile(hex2bin(
            "0aa82c0a0c736368656d612e70726f746f12086d657368706161731a1f67" .
            "6f6f676c652f70726f746f6275662f74696d657374616d702e70726f746f" .
            "1a19676f6f676c652f70726f746f6275662f616e792e70726f746f1a1b67" .
            "6f6f676c652f70726f746f6275662f656d7074792e70726f746f1a366769" .
            "746875622e636f6d2f6d7769746b6f772f676f2d70726f746f2d76616c69" .
            "6461746f72732f76616c696461746f722e70726f746f2293010a09417574" .
            "686e52756c6512220a086a776b735f7572691801200128094210e2df1f0c" .
            "0a0a5e2e7b312c3232357d2412200a066973737565721802200128094210" .
            "e2df1f0c0a0a5e2e7b312c3232357d2412100a0861756469656e63651803" .
            "20032809122e0a146f757075745f7061796c6f61645f6865616465721804" .
            "200128094210e2df1f0c0a0a5e2e7b312c3232357d2422270a0b41757468" .
            "7a536f7572636512180a10616c6c6f775f6e616d65737061636573180120" .
            "03280922520a0c417574687a5375626a65637412150a0d616c6c6f775f69" .
            "73737565727318062003280912130a0b616c6c6f775f726f6c6573180720" .
            "03280912160a0e616c6c6f775f61756469656e636518082003280922680a" .
            "10417574687a44657374696e6174696f6e12130a0b616c6c6f775f706174" .
            "687318022003280912130a0b616c6c6f775f686f73747318032003280912" .
            "150a0d616c6c6f775f6d6574686f647318042003280912130a0b616c6c6f" .
            "775f706f727473180520032809228c010a09417574687a52756c6512250a" .
            "06736f7572636518012001280b32152e6d657368706161732e417574687a" .
            "536f75726365122f0a0b64657374696e6174696f6e18022001280b321a2e" .
            "6d657368706161732e417574687a44657374696e6174696f6e12270a0773" .
            "75626a65637418032001280b32162e6d657368706161732e417574687a53" .
            "75626a656374222b0a05417574687a12220a0572756c657318012003280b" .
            "32132e6d657368706161732e417574687a52756c65222b0a05417574686e" .
            "12220a0572756c657318012003280b32132e6d657368706161732e417574" .
            "686e52756c6522eb010a0b536563726574496e707574121e0a046e616d65" .
            "1801200128094210e2df1f0c0a0a5e2e7b312c3232357d2412210a077072" .
            "6f6a6563741802200128094210e2df1f0c0a0a5e2e7b312c3232357d2412" .
            "220a047479706518032001280e32142e6d657368706161732e5365637265" .
            "745479706512110a09696d6d757461626c6518042001280812350a046461" .
            "746118052003280b321f2e6d657368706161732e536563726574496e7075" .
            "742e44617461456e7472794206e2df1f0220011a2b0a0944617461456e74" .
            "7279120b0a036b6579180120012809120d0a0576616c7565180220012809" .
            "3a02380122e1010a06536563726574121e0a046e616d6518012001280942" .
            "10e2df1f0c0a0a5e2e7b312c3232357d2412210a0770726f6a6563741802" .
            "200128094210e2df1f0c0a0a5e2e7b312c3232357d2412220a0474797065" .
            "18032001280e32142e6d657368706161732e536563726574547970651211" .
            "0a09696d6d757461626c6518042001280812300a04646174611805200328" .
            "0b321a2e6d657368706161732e5365637265742e44617461456e74727942" .
            "06e2df1f0220011a2b0a0944617461456e747279120b0a036b6579180120" .
            "012809120d0a0576616c75651802200128093a02380122d9010a11536572" .
            "766572544c5353657474696e677312160a0e68747470735f726564697265" .
            "6374180120012808121f0a046d6f646518022001280e32112e6d65736870" .
            "6161732e544c536d6f646512170a0f63726564656e7469616c5f6e616d65" .
            "18032001280912190a117375626a6563745f616c745f6e616d6573180420" .
            "032809121f0a177665726966795f63657274696669636174655f73706b69" .
            "180520032809121f0a177665726966795f63657274696669636174655f68" .
            "61736818062003280912150a0d6369706865725f73756974657318072003" .
            "280922c6010a0f476174657761794c697374656e657212140a04706f7274" .
            "18012001280d4206e2df1f021000121e0a046e616d651802200128094210" .
            "e2df1f0c0a0a5e2e7b312c3232357d2412350a0870726f746f636f6c1803" .
            "2001280e321b2e6d657368706161732e5472616e73706f727450726f746f" .
            "636f6c4206e2df1f02100012150a05686f7374731804200328094206e2df" .
            "1f026000122f0a0a746c735f636f6e66696718052001280b321b2e6d6573" .
            "68706161732e536572766572544c5353657474696e677322560a07476174" .
            "65776179120c0a046e616d65180120012809120f0a0770726f6a65637418" .
            "0220012809122c0a096c697374656e65727318032003280b32192e6d6573" .
            "68706161732e476174657761794c697374656e6572227f0a0c4761746577" .
            "6179496e707574121e0a046e616d651801200128094210e2df1f0c0a0a5e" .
            "2e7b312c3232357d2412210a0770726f6a6563741802200128094210e2df" .
            "1f0c0a0a5e2e7b312c3232357d24122c0a096c697374656e657273180320" .
            "03280b32192e6d657368706161732e476174657761794c697374656e6572" .
            "22e3010a0948545450526f757465121e0a046e616d651801200128094210" .
            "e2df1f0c0a0a5e2e7b312c3232357d2412140a04706f727418022001280d" .
            "4206e2df1f02100012130a0b706174685f70726566697818032001280912" .
            "130a0b726577726974655f75726918052001280912150a0d616c6c6f775f" .
            "6f726967696e7318062003280912150a0d616c6c6f775f6d6574686f6473" .
            "18072003280912150a0d616c6c6f775f6865616465727318082003280912" .
            "160a0e6578706f73655f6865616465727318092003280912190a11616c6c" .
            "6f775f63726564656e7469616c73180a2001280822670a0a4e6574776f72" .
            "6b696e6712100a086761746577617973180120032809120d0a05686f7374" .
            "73180220032809120e0a066578706f727418032001280812280a0b687474" .
            "705f726f7574657318042003280b32132e6d657368706161732e48545450" .
            "526f7574652296020a09436f6e7461696e6572121e0a046e616d65180120" .
            "0128094210e2df1f0c0a0a5e2e7b312c3232357d24121f0a05696d616765" .
            "1802200128094210e2df1f0c0a0a5e2e7b312c3232357d24120c0a046172" .
            "677318032003280912290a03656e7618042003280b321c2e6d6573687061" .
            "61732e436f6e7461696e65722e456e76456e74727912350a05706f727473" .
            "18052003280b321e2e6d657368706161732e436f6e7461696e65722e506f" .
            "727473456e7472794206e2df1f0220011a2a0a08456e76456e747279120b" .
            "0a036b6579180120012809120d0a0576616c75651802200128093a023801" .
            "1a2c0a0a506f727473456e747279120b0a036b6579180120012809120d0a" .
            "0576616c756518022001280d3a02380122e6020a03417070121e0a046e61" .
            "6d651801200128094210e2df1f0c0a0a5e2e7b312c3232357d2412210a07" .
            "70726f6a6563741802200128094210e2df1f0c0a0a5e2e7b312c3232357d" .
            "24122f0a0a636f6e7461696e65727318032003280b32132e6d6573687061" .
            "61732e436f6e7461696e65724206e2df1f02200112100a087265706c6963" .
            "617318082001280d12300a0a6e6574776f726b696e67180b2001280b3214" .
            "2e6d657368706161732e4e6574776f726b696e674206e2df1f022001122f" .
            "0a0e61757468656e7469636174696f6e180c2001280b320f2e6d65736870" .
            "6161732e417574686e4206e2df1f022001122e0a0d617574686f72697a61" .
            "74696f6e180d2001280b320f2e6d657368706161732e417574687a4206e2" .
            "df1f02200112190a11696d6167655f70756c6c5f736563726574180e2001" .
            "2809122b0a0673746174757318142001280b32132e6d657368706161732e" .
            "4170705374617475734206e2df1f02200122ce010a045461736b121e0a04" .
            "6e616d651801200128094210e2df1f0c0a0a5e2e7b312c3232357d241221" .
            "0a0770726f6a6563741802200128094210e2df1f0c0a0a5e2e7b312c3232" .
            "357d2412190a11696d6167655f70756c6c5f736563726574180320012809" .
            "122f0a0a636f6e7461696e65727318042003280b32132e6d657368706161" .
            "732e436f6e7461696e65724206e2df1f02200112220a087363686564756c" .
            "651807200128094210e2df1f0c0a0a5e2e7b312c3232357d2412130a0b63" .
            "6f6d706c6574696f6e7318082001280d22d3010a095461736b496e707574" .
            "121e0a046e616d651801200128094210e2df1f0c0a0a5e2e7b312c323235" .
            "7d2412210a0770726f6a6563741802200128094210e2df1f0c0a0a5e2e7b" .
            "312c3232357d2412190a11696d6167655f70756c6c5f7365637265741803" .
            "20012809122f0a0a636f6e7461696e65727318042003280b32132e6d6573" .
            "68706161732e436f6e7461696e65724206e2df1f02200112220a08736368" .
            "6564756c651807200128094210e2df1f0c0a0a5e2e7b312c3232357d2412" .
            "130a0b636f6d706c6574696f6e7318082001280d22b6020a08417070496e" .
            "707574121e0a046e616d651801200128094210e2df1f0c0a0a5e2e7b312c" .
            "3232357d2412210a0770726f6a6563741802200128094210e2df1f0c0a0a" .
            "5e2e7b312c3232357d24122f0a0a636f6e7461696e65727318032003280b" .
            "32132e6d657368706161732e436f6e7461696e65724206e2df1f02200112" .
            "100a087265706c6963617318072001280d12300a0a6e6574776f726b696e" .
            "67180a2001280b32142e6d657368706161732e4e6574776f726b696e6742" .
            "06e2df1f02200112270a0e61757468656e7469636174696f6e180c200128" .
            "0b320f2e6d657368706161732e417574686e122e0a0d617574686f72697a" .
            "6174696f6e180d2001280b320f2e6d657368706161732e417574687a4206" .
            "e2df1f02200112190a11696d6167655f70756c6c5f736563726574180e20" .
            "01280922480a03526566121e0a046e616d651801200128094210e2df1f0c" .
            "0a0a5e2e7b312c3232357d2412210a0770726f6a65637418022001280942" .
            "10e2df1f0c0a0a5e2e7b312c3232357d24223b0a075265706c696361120d" .
            "0a05706861736518012001280912110a09636f6e646974696f6e18022001" .
            "2809120e0a06726561736f6e18032001280922300a094170705374617475" .
            "7312230a087265706c6963617318012003280b32112e6d65736870616173" .
            "2e5265706c69636122160a034c6f67120f0a076d65737361676518012001" .
            "2809222b0a044170707312230a0c6170706c69636174696f6e7318012003" .
            "280b320d2e6d657368706161732e41707022260a055461736b73121d0a05" .
            "7461736b7318012003280b320e2e6d657368706161732e5461736b222e0a" .
            "0c50726f6a656374496e707574121e0a046e616d651801200128094210e2" .
            "df1f0c0a0a5e2e7b312c3232357d2422170a0750726f6a656374120c0a04" .
            "6e616d65180120012809222c0a0a50726f6a656374526566121e0a046e61" .
            "6d651801200128094210e2df1f0c0a0a5e2e7b312c3232357d24221c0a08" .
            "50726f6a6563747312100a0870726f6a656374731801200328092a3d0a0a" .
            "53656372657454797065120a0a064f5041515545100012100a0c544c535f" .
            "434552545f4b4559100112110a0d444f434b45525f434f4e46494710022a" .
            "700a115472616e73706f727450726f746f636f6c12140a10494e56414c49" .
            "445f50524f544f434f4c100012080a0448545450100112090a0548545450" .
            "53100212080a0447525043100312090a054854545032100412090a054d4f" .
            "4e474f100512070a03544350100612070a03544c5310072a5a0a07544c53" .
            "6d6f6465120f0a0b504153535448524f5547481000120a0a0653494d504c" .
            "451001120a0a064d555455414c100212140a104155544f5f504153535448" .
            "524f554748100312100a0c495354494f5f4d555455414c100432b70a0a0f" .
            "4d6573685061617353657276696365123c0a0d43726561746550726f6a65" .
            "637412162e6d657368706161732e50726f6a656374496e7075741a112e6d" .
            "657368706161732e50726f6a6563742200123c0a0c4c69737450726f6a65" .
            "63747312162e676f6f676c652e70726f746f6275662e456d7074791a122e" .
            "6d657368706161732e50726f6a656374732200123f0a0d44656c65746550" .
            "726f6a65637412142e6d657368706161732e50726f6a6563745265661a16" .
            "2e676f6f676c652e70726f746f6275662e456d707479220012370a0a4765" .
            "7450726f6a65637412142e6d657368706161732e50726f6a656374526566" .
            "1a112e6d657368706161732e50726f6a6563742200123c0a0d5570646174" .
            "6550726f6a65637412162e6d657368706161732e50726f6a656374496e70" .
            "75741a112e6d657368706161732e50726f6a656374220012300a09437265" .
            "61746541707012122e6d657368706161732e417070496e7075741a0d2e6d" .
            "657368706161732e417070220012300a0955706461746541707012122e6d" .
            "657368706161732e417070496e7075741a0d2e6d657368706161732e4170" .
            "70220012340a0944656c657465417070120d2e6d657368706161732e5265" .
            "661a162e676f6f676c652e70726f746f6275662e456d707479220012280a" .
            "06476574417070120d2e6d657368706161732e5265661a0d2e6d65736870" .
            "6161732e417070220012320a084c6973744170707312142e6d6573687061" .
            "61732e50726f6a6563745265661a0e2e6d657368706161732e4170707322" .
            "0012330a0a4372656174655461736b12132e6d657368706161732e546173" .
            "6b496e7075741a0e2e6d657368706161732e5461736b220012330a0a5570" .
            "646174655461736b12132e6d657368706161732e5461736b496e7075741a" .
            "0e2e6d657368706161732e5461736b220012350a0a44656c657465546173" .
            "6b120d2e6d657368706161732e5265661a162e676f6f676c652e70726f74" .
            "6f6275662e456d7074792200122a0a074765745461736b120d2e6d657368" .
            "706161732e5265661a0e2e6d657368706161732e5461736b220012340a09" .
            "4c6973745461736b7312142e6d657368706161732e50726f6a6563745265" .
            "661a0f2e6d657368706161732e5461736b732200123c0a0d437265617465" .
            "4761746577617912162e6d657368706161732e47617465776179496e7075" .
            "741a112e6d657368706161732e476174657761792200123c0a0d55706461" .
            "74654761746577617912162e6d657368706161732e47617465776179496e" .
            "7075741a112e6d657368706161732e47617465776179220012380a0d4465" .
            "6c65746547617465776179120d2e6d657368706161732e5265661a162e67" .
            "6f6f676c652e70726f746f6275662e456d707479220012300a0a47657447" .
            "617465776179120d2e6d657368706161732e5265661a112e6d6573687061" .
            "61732e47617465776179220012390a0c4372656174655365637265741215" .
            "2e6d657368706161732e536563726574496e7075741a102e6d6573687061" .
            "61732e536563726574220012390a0c55706461746553656372657412152e" .
            "6d657368706161732e536563726574496e7075741a102e6d657368706161" .
            "732e536563726574220012370a0c44656c657465536563726574120d2e6d" .
            "657368706161732e5265661a162e676f6f676c652e70726f746f6275662e" .
            "456d7074792200122e0a09476574536563726574120d2e6d657368706161" .
            "732e5265661a102e6d657368706161732e5365637265742200122e0a0a53" .
            "747265616d4c6f6773120d2e6d657368706161732e5265661a0d2e6d6573" .
            "68706161732e4c6f6722003001420c5a0a6d657368706161737062620670" .
            "726f746f33"
        ));

        static::$is_initialized = true;
    }
}

