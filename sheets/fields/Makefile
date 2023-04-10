# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make fields.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

fields.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > fields.tex

fields.pdf: ../../*.tex ../../trademark.pdf *.tex fields.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode fields.tex
	../../latexrun fields.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: fields.pdf
	open fields.pdf

o:
	make open

terms: fields.tex
	bbk_sheet -mode ts fields.tex

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
	rm -f fields.*
	make reset
	bbk_sheet -mode mk > Makefile