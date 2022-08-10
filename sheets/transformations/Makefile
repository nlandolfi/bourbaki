# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make transformations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

transformations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > transformations.tex

transformations.pdf: ../../*.tex ../../trademark.pdf *.tex transformations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode transformations.tex
	../../latexrun transformations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: transformations.pdf
	open transformations.pdf

o:
	make open

terms: transformations.tex
	bbk_sheet -mode ts transformations.tex

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
	rm -f transformations.*
	make reset
	bbk_sheet -mode mk > Makefile