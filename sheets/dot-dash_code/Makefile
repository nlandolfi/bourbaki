# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dot-dash_code.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

dot-dash_code.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > dot-dash_code.tex

dot-dash_code.pdf: ../../*.tex ../../trademark.pdf *.tex dot-dash_code.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dot-dash_code.tex
	../../latexrun dot-dash_code.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dot-dash_code.pdf
	open dot-dash_code.pdf

o:
	make open

terms: dot-dash_code.tex
	bbk_sheet -mode ts dot-dash_code.tex

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
	rm -f dot-dash_code.*
	make reset
	bbk_sheet -mode mk > Makefile