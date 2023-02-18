# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make operators.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

operators.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > operators.tex

operators.pdf: ../../*.tex ../../trademark.pdf *.tex operators.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode operators.tex
	../../latexrun operators.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: operators.pdf
	open operators.pdf

o:
	make open

terms: operators.tex
	bbk_sheet -mode ts operators.tex

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
	rm -f operators.*
	make reset
	bbk_sheet -mode mk > Makefile