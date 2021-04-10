# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make quadratic_forms.pdf
	make graph.pdf

quadratic_forms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > quadratic_forms.tex

quadratic_forms.pdf: ../../*.tex ../../trademark.pdf *.tex quadratic_forms.tex
	# pdflatex --file-line-error -interaction=nonstopmode quadratic_forms.tex
	../../latexrun quadratic_forms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quadratic_forms.pdf
	open quadratic_forms.pdf

o:
	make open

terms: quadratic_forms.tex
	bbk_sheet -mode ts quadratic_forms.tex

spell:
	aspell -c quadratic_forms.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f quadratic_forms.*
	make reset
	bbk_sheet -mode mk > Makefile