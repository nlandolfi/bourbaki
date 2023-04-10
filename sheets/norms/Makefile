# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make norms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

norms.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > norms.tex

norms.pdf: ../../*.tex ../../trademark.pdf *.tex norms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode norms.tex
	../../latexrun norms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: norms.pdf
	open norms.pdf

o:
	make open

terms: norms.tex
	bbk_sheet -mode ts norms.tex

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
	rm -f norms.*
	make reset
	bbk_sheet -mode mk > Makefile