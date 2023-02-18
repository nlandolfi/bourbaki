# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make polytopes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

polytopes.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > polytopes.tex

polytopes.pdf: ../../*.tex ../../trademark.pdf *.tex polytopes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode polytopes.tex
	../../latexrun polytopes.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: polytopes.pdf
	open polytopes.pdf

o:
	make open

terms: polytopes.tex
	bbk_sheet -mode ts polytopes.tex

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
	rm -f polytopes.*
	make reset
	bbk_sheet -mode mk > Makefile