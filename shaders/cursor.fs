#version 330

in vec2 fragTexCoord;
in vec4 fragColor;

out vec4 finalColor;

uniform sampler2D tex;

void main() {
    vec4 texColor = texture(tex, fragTexCoord);
    finalColor = texColor;
}
