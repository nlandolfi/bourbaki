# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make basel_problem.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

basel_problem.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > basel_problem.tex

basel_problem.pdf: ../../*.tex ../../trademark.pdf *.tex basel_problem.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode basel_problem.tex
	../../latexrun basel_problem.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: basel_problem.pdf
	open basel_problem.pdf

o:
	make open

terms: basel_problem.tex
	bbk sheet -mode ts basel_problem.tex

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
	rm -f basel_problem.*
	make reset
	bbk sheet -mode mk > Makefile