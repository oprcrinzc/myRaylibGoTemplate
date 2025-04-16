in vec2 fragTexCoord;
in vec4 fragColor;

uniform sampler2d texture0;
uniform vec2 lightSourcePos[10];
uniform int lightSourceNum;

out vec4 finalColor;

void main(){

    finalColor = fragColor;
}