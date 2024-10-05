#version 410 core

in vec2 fragTexCoord;
in vec4 fragColor;
out vec4 finalColor;
uniform sampler2D tex;

uniform int size = 2; // blur size
uniform float separation = 3.0; // spread
uniform float threshold = 0.3;
uniform float amount = 1.0;

void main() {
    vec2 texSize = textureSize(tex, 0).xy;
    float dist = length(fragTexCoord.y - 0.5);

    vec4 result = vec4(0.0);
    float count = 0.0;

    for (int x = -size; x <= size; x++) {
        for (int y = -size; y <= size; y++) {
            vec2 offset = (vec2(x, y) * separation * dist) / texSize;
            vec4 color = texture(tex, fragTexCoord + offset);
            float value = max(color.r, max(color.g, color.b));
            if (value > threshold) {
                color += color * (value - threshold);
            }
            result += color;
            count++;
        }
    }

    finalColor = fragColor * mix(vec4(1.0), result / count, amount);
}
