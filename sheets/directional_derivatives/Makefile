# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make directional_derivatives.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

directional_derivatives.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > directional_derivatives.tex

directional_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex directional_derivatives.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode directional_derivatives.tex
	../../latexrun directional_derivatives.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: directional_derivatives.pdf
	open directional_derivatives.pdf

o:
	make open

terms: directional_derivatives.tex
	bbk_sheet -mode ts directional_derivatives.tex

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
	rm -f directional_derivatives.*
	make reset
	bbk_sheet -mode mk > Makefile