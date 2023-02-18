# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make digital_images.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

digital_images.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > digital_images.tex

digital_images.pdf: ../../*.tex ../../trademark.pdf *.tex digital_images.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode digital_images.tex
	../../latexrun digital_images.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: digital_images.pdf
	open digital_images.pdf

o:
	make open

terms: digital_images.tex
	bbk_sheet -mode ts digital_images.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f digital_images.*
	make reset
	bbk_sheet -mode mk > Makefile