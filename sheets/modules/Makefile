# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make modules.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

modules.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > modules.tex

modules.pdf: ../../*.tex ../../trademark.pdf *.tex modules.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode modules.tex
	../../latexrun modules.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: modules.pdf
	open modules.pdf

o:
	make open

terms: modules.tex
	bbk_sheet -mode ts modules.tex

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
	rm -f modules.*
	make reset
	bbk_sheet -mode mk > Makefile