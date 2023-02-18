# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make finite_automata.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

finite_automata.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > finite_automata.tex

finite_automata.pdf: ../../*.tex ../../trademark.pdf *.tex finite_automata.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode finite_automata.tex
	../../latexrun finite_automata.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: finite_automata.pdf
	open finite_automata.pdf

o:
	make open

terms: finite_automata.tex
	bbk_sheet -mode ts finite_automata.tex

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
	rm -f finite_automata.*
	make reset
	bbk_sheet -mode mk > Makefile